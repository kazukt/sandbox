import { GetServerSideProps, NextPage } from "next";
import { useState } from "react";
import styles from "./index.module.css";

type Props = {
    initialImageUrl: string;
}

const IndexPage: NextPage<Props> = ({ initialImageUrl }) => {
    const [imageUrl, setImageUrl] = useState(initialImageUrl);
    const [loading, setLoading] = useState(false);
    const handleClick = async () => {
        setLoading(true);
        const newImange = await fetchImage();
        setImageUrl(newImange.url);
        setLoading(false);
    };
    return (
        <div className={styles.page}>
            <button onClick={handleClick} className={styles.button}>
                他のニャンコも見る
            </button>
            <div className={styles.frame}>
                {loading || <img src={imageUrl} />}
            </div>
        </div>
    );
};

export default IndexPage;

export const getServerSideProps: GetServerSideProps<Props> = async () => {
    const image = await fetchImage();
    return {
        props: {
            initialImageUrl: image.url,
        },
    };
};

type Image = {
    url: string;
};

const fetchImage = async (): Promise<Image> => {
    const res = await fetch("https://api.thecatapi.com/v1/images/search");
    const images = await res.json();
    console.log(images);
    return images[0];
};
