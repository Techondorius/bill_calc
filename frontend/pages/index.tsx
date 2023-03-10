import { GetServerSideProps } from 'next'

export const getServerSideProps: GetServerSideProps = async () => {
  return {
    redirect: {
      permanent: false, // 永続的なリダイレクトかどうか
      destination: '/login', // リダイレクト先
    },
  }
}

export default getServerSideProps
