using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Confluent.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Confluent;

// Generated from example definition: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Validations_ValidateOrganizations.json
// this example is just showing the usage of "Validations_ValidateOrganization" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
string organizationName = "myOrganization";
ConfluentOrganizationData data = new ConfluentOrganizationData(new AzureLocation("West US"), new ConfluentOfferDetail("string", "string", "string", "string", "string")
{
    PrivateOfferId = "string",
    PrivateOfferIds = { "string" },
}, new ConfluentUserDetail("abc@microsoft.com")
{
    FirstName = "string",
    LastName = "string",
    UserPrincipalName = "abc@microsoft.com",
    AadEmail = "abc@microsoft.com",
})
{
    Tags =
    {
    ["Environment"] = "Dev"
    },
};
ConfluentOrganizationResource result = await resourceGroupResource.ValidateOrganizationAsync(organizationName, data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ConfluentOrganizationData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
