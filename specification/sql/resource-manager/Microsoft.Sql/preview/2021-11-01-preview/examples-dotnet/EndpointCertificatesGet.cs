using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2021-11-01-preview/examples/EndpointCertificatesGet.json
// this example is just showing the usage of "EndpointCertificates_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedInstanceResource created on azure
// for more information of creating ManagedInstanceResource, please refer to the document of ManagedInstanceResource
string subscriptionId = "38e0dc56-907f-45ba-a97c-74233baad471";
string resourceGroupName = "testrg";
string managedInstanceName = "testcl";
ResourceIdentifier managedInstanceResourceId = ManagedInstanceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName);
ManagedInstanceResource managedInstance = client.GetManagedInstanceResource(managedInstanceResourceId);

// get the collection of this EndpointCertificateResource
EndpointCertificateCollection collection = managedInstance.GetEndpointCertificates();

// invoke the operation
string endpointType = "DATABASE_MIRRORING";
NullableResponse<EndpointCertificateResource> response = await collection.GetIfExistsAsync(endpointType);
EndpointCertificateResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EndpointCertificateData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
