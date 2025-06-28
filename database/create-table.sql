CREATE TABLE public.users
(
    id uuid DEFAULT gen_random_uuid(),
    email text NOT NULL,
    password_hash text NOT NULL,
    name text,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id),
    UNIQUE (email)
);

ALTER TABLE IF EXISTS public.users
    OWNER to postgres;

  
CREATE TABLE public.artists
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    country TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)
;

ALTER TABLE IF EXISTS public.artists
    OWNER to postgres;

CREATE TABLE public.genres
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
)
;

ALTER TABLE IF EXISTS public.genres
    OWNER to postgres;


CREATE TABLE public.vinyls
(
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  title text NOT NULL,
  price numeric(10,2) NOT NULL,
  stock_count integer NOT NULL DEFAULT 0,
  cover_url text,
  artist_id uuid NOT NULL,
  genre_id uuid NOT NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),

  CONSTRAINT fk_artist FOREIGN KEY (artist_id)
        REFERENCES public.artists (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID,
    CONSTRAINT fk_genre FOREIGN KEY (genre_id)
        REFERENCES public.genres (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)
;

ALTER TABLE IF EXISTS public.vinyls
    OWNER to postgres;

CREATE TABLE public.orders
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    total_amount NUMERIC(10,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);

ALTER TABLE IF EXISTS public.orders
    OWNER to postgres;


CREATE TABLE public.order_items
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id UUID NOT NULL,
    vinyl_id UUID NOT NULL,
    quantity INTEGER NOT NULL DEFAULT 1,
    unit_price NUMERIC(10,2) NOT NULL,

    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id),
    CONSTRAINT fk_vinyl FOREIGN KEY (vinyl_id) REFERENCES vinyls(id)
);

ALTER TABLE IF EXISTS public.order_items
    OWNER to postgres;


CREATE TABLE public.news
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE IF EXISTS public.news
    OWNER to postgres;


CREATE TABLE public.admins
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    name TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

ALTER TABLE IF EXISTS public.admins
    OWNER to postgres;
