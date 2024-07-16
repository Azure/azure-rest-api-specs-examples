const { InformaticaDataManagement } = require("@azure/arm-informaticadatamanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a InformaticaOrganizationResource
 *
 * @summary Update a InformaticaOrganizationResource
 * x-ms-original-file: specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/Organizations_Update_MaximumSet_Gen.json
 */
async function organizationsUpdate() {
  const subscriptionId =
    process.env["INFORMATICA_SUBSCRIPTION_ID"] || "3599DA28-E346-4D9F-811E-189C0445F0FE";
  const resourceGroupName = process.env["INFORMATICA_RESOURCE_GROUP"] || "rgopenapi";
  const organizationName = "_-";
  const properties = {
    properties: {
      companyDetails: {
        business: "mwqblnruflwpolgbxpqbqneve",
        companyName: "xkrvbozrjcvappqeeyt",
        country: "rvlzppgvopcw",
        domain: "dponvwnrdrnzahcurqssesukbsokdd",
        numberOfEmployees: 22,
        officeAddress: "sfcx",
      },
      existingResourceId:
        "/subscriptions/subid/resourceGroups/rg1/providers/Informatica.DataManagement/organizations/org1/serverlessRuntimes/serverlessRuntimeName",
      marketplaceDetails: {
        marketplaceSubscriptionId: "szhyxzgjtssjmlguivepc",
        offerDetails: {
          offerId: "idaxbflabvjsippplyenvrpgeydsjxcmyubgukffkcdvlvrtwpdhnvdblxjsldiuswrchsibk",
          planId: "giihvvnwdwzkfqrhkpqzbgfotzyixnsvmxzauseebillhslauglzfxzvzvts",
          planName:
            "tfqjenotaewzdeerliteqxdawuqxhwdzbtiiimsaedrlsnbdoonnloakjtvnwhhrcyxxsgoachguthqvlahpjyofpoqpfacfmiaauawazkmxkjgvktbptojknzojtjrfzvbbjjkvstabqyaczxinijhoxrjukftsagpwgsvpmczopztmplipyufhuaumfx",
          publisherId:
            "ktzfghsyjqbsswhltoaemgotmnorhdogvkaxplutbjjqzuepxizliynyakersobagvpwvpzwjtjjxigsqgcyqaahaxdijghnexliofhfjlqzjmmbvrhcvjxdodnexxizbgfhjopbwzjojxsluasnwwsgcajefglbcvzpaeblanhmurcculndtfwnfjyxol",
          termId: "eolmwogtgpdncqoigqcdomupwummaicwvdxgbskpdsmjizdfbdgbxbuekcpwmenqzbhqxpdnjtup",
          termUnit: "nykqoplazujcwmfldntifjqrnx",
        },
      },
      userDetails: {
        emailAddress: "7_-46@13D--3.m-4x-.11.c-9-.DHLYFc",
        firstName: "qguqrmanyupoi",
        lastName: "ugzg",
        phoneNumber: "uxa",
        upn: "viwjrkn",
      },
    },
    tags: { key1918: "fbjvtuvzsghpl" },
  };
  const credential = new DefaultAzureCredential();
  const client = new InformaticaDataManagement(credential, subscriptionId);
  const result = await client.organizations.update(resourceGroupName, organizationName, properties);
  console.log(result);
}
