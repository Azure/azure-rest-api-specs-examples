const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Creates or updates a domain.
 *
 * @summary Description for Creates or updates a domain.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-11-01/examples/CreateAppServiceDomain.json
 */
async function createAppServiceDomain() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const domainName = "example.com";
  const domain = {
    authCode: "exampleAuthCode",
    autoRenew: true,
    consent: {
      agreedAt: new Date("2021-09-10T19:30:53Z"),
      agreedBy: "192.0.2.1",
      agreementKeys: ["agreementKey1"],
    },
    contactAdmin: {
      addressMailing: {
        address1: "3400 State St",
        city: "Chicago",
        country: "United States",
        postalCode: "67098",
        state: "IL",
      },
      email: "admin@email.com",
      fax: "1-245-534-2242",
      jobTitle: "Admin",
      nameFirst: "John",
      nameLast: "Doe",
      nameMiddle: "",
      organization: "Microsoft Inc.",
      phone: "1-245-534-2242",
    },
    contactBilling: {
      addressMailing: {
        address1: "3400 State St",
        city: "Chicago",
        country: "United States",
        postalCode: "67098",
        state: "IL",
      },
      email: "billing@email.com",
      fax: "1-245-534-2242",
      jobTitle: "Billing",
      nameFirst: "John",
      nameLast: "Doe",
      nameMiddle: "",
      organization: "Microsoft Inc.",
      phone: "1-245-534-2242",
    },
    contactRegistrant: {
      addressMailing: {
        address1: "3400 State St",
        city: "Chicago",
        country: "United States",
        postalCode: "67098",
        state: "IL",
      },
      email: "registrant@email.com",
      fax: "1-245-534-2242",
      jobTitle: "Registrant",
      nameFirst: "John",
      nameLast: "Doe",
      nameMiddle: "",
      organization: "Microsoft Inc.",
      phone: "1-245-534-2242",
    },
    contactTech: {
      addressMailing: {
        address1: "3400 State St",
        city: "Chicago",
        country: "United States",
        postalCode: "67098",
        state: "IL",
      },
      email: "tech@email.com",
      fax: "1-245-534-2242",
      jobTitle: "Tech",
      nameFirst: "John",
      nameLast: "Doe",
      nameMiddle: "",
      organization: "Microsoft Inc.",
      phone: "1-245-534-2242",
    },
    dnsType: "DefaultDomainRegistrarDns",
    location: "global",
    privacy: false,
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.domains.beginCreateOrUpdateAndWait(
    resourceGroupName,
    domainName,
    domain,
  );
  console.log(result);
}
