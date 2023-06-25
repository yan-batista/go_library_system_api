"use client";
import styles from "./table.module.scss";
import { useEffect, useState } from "react";
import Image from "next/image";

import ArrowUp from "../../public/arrow_up.svg";
import ArrowDown from "../../public/arrow_down.svg";

type Props = {
  data: { id: number; name: string; author: string; slug: string; publisher: string; isbn: string; quantity: number }[];
  rowsPerPage: number;
};

const BooksTable: React.FC<Props> = ({ data, rowsPerPage }: Props) => {
  const [filtersOpen, setFiltersOpen] = useState(false);
  const [page, setPage] = useState(1);
  const [range, setRange] = useState<number[]>([]);
  const [slicedData, setSlicedData] = useState<Props["data"]>([
    { id: 0, name: "", author: "", slug: "", publisher: "", isbn: "", quantity: 0 },
  ]);

  function changePage(page_number: number) {
    setPage(page_number);
  }

  function movePreviousPage() {
    setPage((prevState) => prevState - 1);
  }

  function moveNextPage() {
    setPage((prevState) => prevState + 1);
  }

  function calculatePageRange(data_length: number, rowsPerPage: number) {
    const range = [];
    const num = Math.ceil(data_length / rowsPerPage);
    let i = 1;
    for (let i = 1; i <= num; i++) {
      range.push(i);
    }
    return range;
  }

  function handleFiltersClick() {
    setFiltersOpen((prevState) => !prevState);
  }

  useEffect(() => {
    const range = calculatePageRange(data.length, rowsPerPage);
    setRange(range);

    const newData = data.slice((page - 1) * rowsPerPage, page * rowsPerPage);
    setSlicedData([...newData]);
  }, [page]);

  return (
    <div className={`${styles.center} ${styles.column}`}>
      <div className={`${styles.row} ${styles.table_filter_title}`} onClick={handleFiltersClick}>
        <h2 className={styles.prevent_select}>Filters</h2>
        <Image alt="show or hide filters" src={filtersOpen ? ArrowUp : ArrowDown} />
      </div>
      {filtersOpen && (
        <div className={`${styles.table_filter} ${styles.column}`}>
          <div>
            <label htmlFor="name">Name: </label>
            <input name="name" placeholder="Book Name" />
          </div>

          <div>
            <label htmlFor="author">Author: </label>
            <input name="author" placeholder="Author Name" />
          </div>

          <div className={`${styles.row} ${styles.publisher_isbn_row}`}>
            <div>
              <label htmlFor="publisher">Publisher: </label>
              <input name="publisher" placeholder="Publisher Name" />
            </div>

            <div>
              <label htmlFor="isbn">ISBN: </label>
              <input type="text" name="isbn" placeholder="000-0000000000" />
            </div>
          </div>
        </div>
      )}
      <div className={`${styles.table_header} ${styles.row}`}>
        {page != range[0] && <p onClick={() => movePreviousPage()}> {"<"} </p>}
        {range.map((page_number) => {
          return (
            <p
              key={page_number}
              className={`${page_number === page ? styles.selected_page : ""}`}
              onClick={() => changePage(page_number)}
            >
              {page_number}
            </p>
          );
        })}
        {page != range[range.length - 1] && <p onClick={() => moveNextPage()}> {">"} </p>}
      </div>
      <table className={`${styles.table}`}>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Author</th>
            <th>Slug</th>
            <th>Publisher</th>
            <th>ISBN</th>
            <th>Quantity</th>
          </tr>
        </thead>
        <tbody>
          {slicedData.map((row) => {
            return (
              <tr key={row.id}>
                <td>{row.id}</td>
                <td>{row.name}</td>
                <td>{row.author}</td>
                <td>{row.slug}</td>
                <td>{row.publisher}</td>
                <td>{row.isbn}</td>
                <td>{row.quantity}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
};

export default BooksTable;
