import JsEncrypt from 'jsencrypt';
import Config from '@/config/index'

const PublicKey = Config.PUBLIC_KEY;
/**
 * rsa算法公钥加密
 * @param params
 */
// eslint-disable-next-line camelcase, no-undef
export function rsaEncrypt(params: object | string): string | false {
  const Encrypt = new JsEncrypt();

  Encrypt.setPublicKey(PublicKey);

  const date = new Date();
  const dateYear = date.getFullYear();
  const dateMonth = (date.getMonth() + 1).toString().padStart(2, '0');
  const dateDate = date.getDate().toString().padStart(2, '0');
  const dateHours = date.getHours().toString().padStart(2, '0');
  const dateMinutes = date.getMinutes().toString().padStart(2, '0');
  const str = '' + dateYear + dateMonth + dateDate + dateHours + dateMinutes

  const timestamp = "" + parseInt(str);
  // console.log(timestamp)

  const param = params.toString() + timestamp;
  const data = Encrypt.encrypt(
    typeof params === 'object' ? JSON.stringify(params) : param
  );
  return data
}
