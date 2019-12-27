import axios from 'axios';

const baseDomain = process.env.VUE_APP_GO_API;
const baseURL = `${baseDomain}/api`;

export default axios.create({
    baseURL,
});
