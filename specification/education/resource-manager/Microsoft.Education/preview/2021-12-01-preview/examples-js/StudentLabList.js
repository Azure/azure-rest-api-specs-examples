const { EducationManagementClient } = require("@azure/arm-education");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all labs associated with the caller of the API.
 *
 * @summary Get a list of all labs associated with the caller of the API.
 * x-ms-original-file: specification/education/resource-manager/Microsoft.Education/preview/2021-12-01-preview/examples/StudentLabList.json
 */
async function studentLabList() {
  const credential = new DefaultAzureCredential();
  const client = new EducationManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.studentLabs.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

studentLabList().catch(console.error);
