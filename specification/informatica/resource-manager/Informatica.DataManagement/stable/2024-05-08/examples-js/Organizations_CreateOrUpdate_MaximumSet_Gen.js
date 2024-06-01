const { InformaticaDataManagement } = require("@azure/arm-informaticadatamanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a InformaticaOrganizationResource
 *
 * @summary Create a InformaticaOrganizationResource
 * x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_CreateOrUpdate_MaximumSet_Gen.json
 */
async function organizationsCreateOrUpdate() {
  const subscriptionId =
    process.env["INFORMATICA_SUBSCRIPTION_ID"] || "3599DA28-E346-4D9F-811E-189C0445F0FE";
  const resourceGroupName = process.env["INFORMATICA_RESOURCE_GROUP"] || "rgopenapi";
  const organizationName = "C";
  const resource = {
    location: "pamjoudtssthlbhrnfjidr",
    properties: {
      companyDetails: {
        business: "pucosrtjv",
        companyName: "xszcggknokhw",
        country: "gwkcpnwyaqc",
        domain: "utcxetzzpmbvwmjrvphqngvp",
        numberOfEmployees: 25,
        officeAddress: "sbttzwyajgdbsvipuiclbzvkcvwyil",
      },
      informaticaProperties: {
        informaticaRegion: "zfqodqpbeflhedypiijdkc",
        organizationId: "wtdmhlwhkvgqdumaehgfgiqcxgnqpx",
        organizationName: "nomzbvwe",
        singleSignOnUrl: "espcbzjhtmgwfdkckhqk",
      },
      linkOrganization: { token: "jjfouhoqpumjvrdsfbimgcy" },
      marketplaceDetails: {
        marketplaceSubscriptionId: "ovenlecocg",
        offerDetails: {
          offerId:
            "cwswcfwmzhjcoksmueukegwaptvpcmbfyvixfhvgwnjyblqivqdkkwkunkgimiopwwkvgnwclmajhuty",
          planId: "jfnemevyivtlxhectiutdavdgfyidolivuojumdzckp",
          planName:
            "iaoxgaitteuoqgujkgxbdgryaobtkjjecuvchwutntrvmuorikrbqqegmelenbewhakiysprrnovjixyxrikscaptrbapbdspu",
          publisherId:
            "zajxpfacudwongxjvnnuhhpygmnydchgowjccyuzsjonegmqxcqqpnzafanggowfqdixnnutyfvmvwrkx",
          termId:
            "tcvvsxdjnjlfmjhmvwklptdmxetnzydxyuhfqchoubmtoeqbchnfxoxqzezlgpxdnzyvzgkynjxzzgetkqccxvpzahxattluqdipvbdktqmndfefitzuifqjpschzlbvixnvznkmmgjwvkplfhemnapsewgqxggdzdokryhv",
          termUnit: "gjwmgevrblbosuogsvfspsgspetbnxaygkbelvadpgwiywl",
        },
      },
      provisioningState: "Accepted",
      userDetails: {
        emailAddress: "7_-46@13D--3.m-4x-.11.c-9-.DHLYFc",
        firstName: "appvdclawzfjntdfdftjevlhvzropnxqtnypid",
        lastName: "nzirbvzmkxtbrlamyatlcszebxgcyncxoascojsmacwvjsjvn",
        phoneNumber: "fvcjylxlmhdnshsgywnzlyvshu",
        upn: "undljch",
      },
    },
    tags: { key8430: "cagshqtjlxtqqhdwtchokvxszybp" },
  };
  const credential = new DefaultAzureCredential();
  const client = new InformaticaDataManagement(credential, subscriptionId);
  const result = await client.organizations.beginCreateOrUpdateAndWait(
    resourceGroupName,
    organizationName,
    resource,
  );
  console.log(result);
}
