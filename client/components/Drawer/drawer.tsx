"use client";

import { useState } from "react";
import { usePathname } from "next/navigation";
import styles from "./drawer.module.scss";
import Link from "next/link";
import Image from "next/image";

// Icons
import Clients from "../../public/clients.svg";
import Book from "../../public/book.svg";
import RentedBook from "../../public/rented_book.svg";
import Settings from "../../public/settings.svg";
import Docs from "../../public/docs.svg";
import ArrowRight from "../../public/arrow_right.svg";
import Close from "../../public/close.svg";
import Logo from "../../public/logo.svg";

const Drawer = () => {
  const [open, setOpen] = useState(false);
  const pathname = usePathname();

  const openCloseDrawer = () => {
    setOpen((prevState) => !prevState);
  };

  return (
    <div className={`${styles.drawer_container} ${open ? styles.drawer_open : styles.drawer_closed}`}>
      <div className={styles.drawer_container_close_button}>
        <p onClick={openCloseDrawer}>
          {open ? (
            <Image alt="icon to minimize drawer side menu" src={Close} />
          ) : (
            <Image alt="icon to expand drawer side menu" src={ArrowRight} />
          )}
        </p>
      </div>

      <div className={styles.drawer_container_title}>
        <h1>
          {open ? (
            <Link href="/">
              <Image alt="logo" src={Logo} className="logo" /> Lib
            </Link>
          ) : (
            <Link href="/">
              <Image alt="logo" src={Logo} className="logo" />
            </Link>
          )}
        </h1>
      </div>

      <nav>
        <ul className={styles.menu_list}>
          <li className={pathname.startsWith("/catalog") ? styles.isActive : ""}>
            <Link href="/catalog">
              <Image alt="book icon for book catalog" src={Book} /> {open ? "Catalog" : ""}
            </Link>
          </li>
          <li className={pathname.startsWith("/clients") ? styles.isActive : ""}>
            <Link href="/clients">
              <Image alt="clients icon for clients list" src={Clients} /> {open ? "Clients" : ""}
            </Link>
          </li>
          <li className={pathname.startsWith("/rented_books") ? styles.isActive : ""}>
            <Link href="rented_books">
              <Image alt="rented books icon for rented book list" src={RentedBook} /> {open ? "Rented Books" : ""}
            </Link>
          </li>
        </ul>

        <ul className={styles.menu_list}>
          <li className={pathname.startsWith("/settings") ? styles.isActive : ""}>
            <Link href="/settings">
              <Image alt="settings icon for settings list" src={Settings} /> {open ? "Settings" : ""}
            </Link>
          </li>
          <li className={pathname.startsWith("/docs") ? styles.isActive : ""}>
            <Link href="/docs">
              <Image alt="docs icon for api docs page" src={Docs} /> {open ? "Docs" : ""}
            </Link>
          </li>
        </ul>
      </nav>
    </div>
  );
};

export default Drawer;
