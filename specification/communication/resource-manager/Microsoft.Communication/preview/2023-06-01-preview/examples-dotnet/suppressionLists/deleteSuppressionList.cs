using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/preview/2023-06-01-preview/examples/suppressionLists/deleteSuppressionList.json
// this example is just showing the usage of "SuppressionLists_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SuppressionListResource created on azure
// for more information of creating SuppressionListResource, please refer to the document of SuppressionListResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "MyResourceGroup";
string emailServiceName = "MyEmailServiceResource";
string domainName = "mydomain.com";
string suppressionListName = "aaaa1111-bbbb-2222-3333-aaaa11112222";
ResourceIdentifier suppressionListResourceId = SuppressionListResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, emailServiceName, domainName, suppressionListName);
SuppressionListResource suppressionListResource = client.GetSuppressionListResource(suppressionListResourceId);

// invoke the operation
await suppressionListResource.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
