using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HealthcareApis;
using Azure.ResourceManager.HealthcareApis.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/legacy/ServiceCreateMinimum.json
// this example is just showing the usage of "Services_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this HealthcareApisServiceResource
HealthcareApisServiceCollection collection = resourceGroupResource.GetHealthcareApisServices();

// invoke the operation
string resourceName = "service2";
HealthcareApisServiceData data = new HealthcareApisServiceData(new AzureLocation("westus2"), HealthcareApisKind.FhirR4)
{
    Properties = new HealthcareApisServiceProperties()
    {
        AccessPolicies =
        {
        new HealthcareApisServiceAccessPolicyEntry("c487e7d1-3210-41a3-8ccc-e9372b78da47")
        },
    },
    Tags =
    {
    },
};
ArmOperation<HealthcareApisServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
HealthcareApisServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
HealthcareApisServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
