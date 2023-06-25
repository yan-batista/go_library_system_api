import BooksTable from "../../components/BooksTable/table";
import axios from "axios";
import styles from "./page.module.scss";

export default async function () {
  const res = await axios.get("http://localhost:3001/api/books").then((res) => res);
  return (
    <main className={`${styles.main}`}>
      <h1>Catalog</h1>
      <BooksTable data={res.data} rowsPerPage={10} />
    </main>
  );
}
