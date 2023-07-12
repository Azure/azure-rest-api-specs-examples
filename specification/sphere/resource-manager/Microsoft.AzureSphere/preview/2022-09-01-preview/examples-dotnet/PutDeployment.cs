using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sphere;

// Generated from example definition: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PutDeployment.json
// this example is just showing the usage of "Deployments_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SphereDeploymentResource created on azure
// for more information of creating SphereDeploymentResource, please refer to the document of SphereDeploymentResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "MyResourceGroup1";
string catalogName = "MyCatalog1";
string productName = "MyProduct1";
string deviceGroupName = "myDeviceGroup1";
string deploymentName = "MyDeployment1";
ResourceIdentifier sphereDeploymentResourceId = SphereDeploymentResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, catalogName, productName, deviceGroupName, deploymentName);
SphereDeploymentResource sphereDeployment = client.GetSphereDeploymentResource(sphereDeploymentResourceId);

// invoke the operation
SphereDeploymentData data = new SphereDeploymentData();
ArmOperation<SphereDeploymentResource> lro = await sphereDeployment.UpdateAsync(WaitUntil.Completed, data);
SphereDeploymentResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SphereDeploymentData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
