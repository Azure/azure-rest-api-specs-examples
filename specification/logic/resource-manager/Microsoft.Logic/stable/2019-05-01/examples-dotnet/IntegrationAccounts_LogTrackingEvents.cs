using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_LogTrackingEvents.json
// this example is just showing the usage of "IntegrationAccounts_LogTrackingEvents" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IntegrationAccountResource created on azure
// for more information of creating IntegrationAccountResource, please refer to the document of IntegrationAccountResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string integrationAccountName = "testIntegrationAccount";
ResourceIdentifier integrationAccountResourceId = IntegrationAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, integrationAccountName);
IntegrationAccountResource integrationAccount = client.GetIntegrationAccountResource(integrationAccountResourceId);

// invoke the operation
IntegrationAccountTrackingEventsContent content = new IntegrationAccountTrackingEventsContent("Microsoft.Logic/workflows", new IntegrationAccountTrackingEvent[]
{
new IntegrationAccountTrackingEvent(IntegrationAccountEventLevel.Informational,DateTimeOffset.Parse("2016-08-05T01:54:49.505567Z"),IntegrationAccountTrackingRecordType.AS2Message)
{
Record = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
{
["agreementProperties"] = new Dictionary<string, object>()
{
["agreementName"] = "testAgreement",
["as2From"] = "testas2from",
["as2To"] = "testas2to",
["receiverPartnerName"] = "testPartner2",
["senderPartnerName"] = "testPartner1"},
["messageProperties"] = new Dictionary<string, object>()
{
["IsMessageEncrypted"] = "false",
["IsMessageSigned"] = "false",
["correlationMessageId"] = "Unique message identifier",
["direction"] = "Receive",
["dispositionType"] = "received-success",
["fileName"] = "test",
["isMdnExpected"] = "true",
["isMessageCompressed"] = "false",
["isMessageFailed"] = "false",
["isNrrEnabled"] = "true",
["mdnType"] = "Async",
["messageId"] = "12345"}}),
Error = new IntegrationAccountTrackingEventErrorInfo()
{
Message = "Some error occurred",
Code = "NotFound",
},
}
});
await integrationAccount.LogTrackingEventsAsync(content);

Console.WriteLine($"Succeeded");
