using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Communication;

// Generated from example definition: specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/suppressionLists/getAddress.json
// this example is just showing the usage of "SuppressionListAddresses_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EmailSuppressionListResource created on azure
// for more information of creating EmailSuppressionListResource, please refer to the document of EmailSuppressionListResource
string subscriptionId = "11112222-3333-4444-5555-666677778888";
string resourceGroupName = "contosoResourceGroup";
string emailServiceName = "contosoEmailService";
string domainName = "contoso.com";
string suppressionListName = "aaaa1111-bbbb-2222-3333-aaaa11112222";
ResourceIdentifier emailSuppressionListResourceId = EmailSuppressionListResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, emailServiceName, domainName, suppressionListName);
EmailSuppressionListResource emailSuppressionList = client.GetEmailSuppressionListResource(emailSuppressionListResourceId);

// get the collection of this EmailSuppressionListAddressResource
EmailSuppressionListAddressCollection collection = emailSuppressionList.GetEmailSuppressionListAddresses();

// invoke the operation
string addressId = "11112222-3333-4444-5555-aaaabbbbcccc";
NullableResponse<EmailSuppressionListAddressResource> response = await collection.GetIfExistsAsync(addressId);
EmailSuppressionListAddressResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EmailSuppressionListAddressData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
