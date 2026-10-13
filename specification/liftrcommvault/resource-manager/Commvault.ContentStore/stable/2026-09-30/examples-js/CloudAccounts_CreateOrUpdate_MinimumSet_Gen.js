const { ContentStoreClient } = require("@azure/arm-commvaultcontentstore");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a CloudAccount
 *
 * @summary create a CloudAccount
 * x-ms-original-file: 2026-09-30/CloudAccounts_CreateOrUpdate_MinimumSet_Gen.json
 */
async function cloudAccountsCreateOrUpdateMinimumSetCCACreateWithRoleAssignments() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "65D4E6D7-7063-4C4B-BAC5-13C45474009E";
  const client = new ContentStoreClient(credential, subscriptionId);
  const result = await client.cloudAccounts.createOrUpdate(
    "rgcommvault",
    "sample-cloudAccountName",
    {
      properties: {
        marketplace: {
          subscriptionId: "tblwyuznrazgchhfczgtlaifwamndt",
          offerDetails: {
            publisherId: "npghpdbgiohslbbeihxdwucejb",
            offerId: "recofyvhkddgkuvducosjstenmy",
            planId: "pqoyqqavjh",
            planName: "hwcltkdvndwfmmnthzwvocujri",
            termUnit: "wzrzqyfzrpqhy",
            termId: "avpgkctrkwdmudsz",
          },
        },
        user: {
          firstName: "John",
          lastName: "Doe",
          emailAddress: "john.doe@contoso.com",
          upn: "john.doe@contoso.com",
          phoneNumber: "1234567890",
        },
        roleAssignmentsOnCcaCreate: [
          {
            roleName: "BackupAdmin",
            entities: [
              {
                id: "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
                displayName: "Tenant Admins",
                entityType: "Group",
              },
            ],
          },
          {
            roleName: "BackupUser",
            entities: [
              {
                id: "22222222-3333-4444-5555-666666666666",
                displayName: "Backup Users SG",
                entityType: "Group",
              },
              {
                id: "33333333-4444-5555-6666-777777777777",
                displayName: "Jane Doe",
                entityType: "User",
              },
            ],
          },
          {
            roleName: "BackupOperator",
            entities: [
              {
                id: "44444444-5555-6666-7777-888888888888",
                displayName: "Ops Team",
                entityType: "Group",
              },
            ],
          },
          {
            roleName: "MultiPersonAuthorization",
            entities: [
              {
                id: "11111111-2222-3333-4444-555555555555",
                displayName: "MPA Approvers",
                entityType: "Group",
              },
            ],
          },
        ],
      },
      identity: { type: "None", userAssignedIdentities: {} },
      tags: {},
      location: "eastus",
    },
  );
  console.log(result);
}
