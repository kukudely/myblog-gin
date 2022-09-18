import{_ as y,r as i,o as c,c as d,b as o,w as l,e as t,F as _,t as a,f as p,a as C}from"./index.06bf44e4.js";const b={props:["id"],data(){return{inputName:"",inputEmail:"",textareaComment:"",artInfo:{},commentList:[],comment:{content:""},total:0,headers:{username:"",user_id:0},queryParam:{pagesize:5,pagenum:1}}},created(){this.getArtInfo(),this.getCommentList()},methods:{async getArtInfo(){const{data:e}=await this.$http.get(`article/info/${this.id}`);this.artInfo=e.data,window.sessionStorage.setItem("title",this.artInfo.title)},async getCommentList(){const{data:e}=await this.$http.get(`commentfront/${this.id}`,{params:{pagesize:this.queryParam.pagesize,pagenum:this.queryParam.pagenum}});this.commentList=e.data,this.total=e.total},async pushComment(){const{data:e}=await this.$http.post("addcomment",{article_id:parseInt(this.id),content:this.textareaComment,username:this.inputName,status:1});if(e.status!==200)return this.$message.error(e.message);this.$message.success("\u8BC4\u8BBA\u6210\u529F\uFF0C\u5F85\u5BA1\u6838\u540E\u663E\u793A"),this.getCommentList()}}},v={style:{"font-size":"25px",padding:"5px","font-family":"\u5FAE\u8F6F\u96C5\u9ED1"}},I={style:{padding:"5px","font-family":"\u5FAE\u8F6F\u96C5\u9ED1"}},w=p("\u63D0\u4EA4"),V=p("\u8BC4\u8BBA\u6570:"),L={style:{padding:"5px",border:"1px","border-color":"gainsboro","border-style":"solid","border-radius":"4px"}};function N(e,r,k,z,s,h){const f=i("el-image"),m=i("el-card"),u=i("el-input"),g=i("el-button"),x=i("font");return c(),d(_,null,[o(m,{class:"box-card box-card-main"},{default:l(()=>[t("div",null,[o(f,{src:s.artInfo.img},null,8,["src"])]),t("div",v,a(s.artInfo.title),1),t("div",I," \xA0\xA0\xA0\xA0"+a(s.artInfo.content),1)]),_:1}),t("div",null,[o(m,{class:"box-card",style:{"margin-top":"20px"}},{default:l(()=>[o(u,{class:"comment",modelValue:s.inputName,"onUpdate:modelValue":r[0]||(r[0]=n=>s.inputName=n),placeholder:"\u8BF7\u8F93\u5165\u540D\u5B57"},null,8,["modelValue"]),o(u,{class:"comment",modelValue:s.textareaComment,"onUpdate:modelValue":r[1]||(r[1]=n=>s.textareaComment=n),autosize:{minRows:2,maxRows:4},type:"textarea",placeholder:"\u8BF7\u8F93\u5165\u5185\u5BB9"},null,8,["modelValue"]),o(g,{class:"comment",onClick:h.pushComment},{default:l(()=>[w]),_:1},8,["onClick"])]),_:1})]),t("div",null,[o(m,{class:"box-card",style:{"margin-top":"20px"}},{default:l(()=>[t("div",null,[o(x,null,{default:l(()=>[V]),_:1}),p(" "+a(this.total),1)]),(c(!0),d(_,null,C(s.commentList,n=>(c(),d("div",{key:n.id,class:"text item",style:{padding:"5px"}},[t("div",L,[t("p",null,a(n.username),1),t("p",null,a(n.CreatedAt),1),t("p",null,a(n.content),1)])]))),128))]),_:1})])],64)}const q=y(b,[["render",N]]);export{q as default};
