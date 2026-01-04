using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/deleteAddress.json
// this example is just showing the usage of "SuppressionListAddresses_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EmailSuppressionListAddressResource created on azure
// for more information of creating EmailSuppressionListAddressResource, please refer to the document of EmailSuppressionListAddressResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "MyResourceGroup";
string emailServiceName = "MyEmailServiceResource";
string domainName = "mydomain.com";
string suppressionListName = "aaaa1111-bbbb-2222-3333-aaaa11112222";
string addressId = "11112222-3333-4444-5555-999999999999";
ResourceIdentifier emailSuppressionListAddressResourceId = EmailSuppressionListAddressResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, emailServiceName, domainName, suppressionListName, addressId);
EmailSuppressionListAddressResource emailSuppressionListAddress = client.GetEmailSuppressionListAddressResource(emailSuppressionListAddressResourceId);

// invoke the operation
await emailSuppressionListAddress.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
