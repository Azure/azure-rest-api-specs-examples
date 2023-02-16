using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeploymentManager;
using Azure.ResourceManager.DeploymentManager.Models;

// Generated from example definition: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate.json
// this example is just showing the usage of "ServiceUnits_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceResource created on azure
// for more information of creating ServiceResource, please refer to the document of ServiceResource
string subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
string resourceGroupName = "myResourceGroup";
string serviceTopologyName = "myTopology";
string serviceName = "myService";
ResourceIdentifier serviceResourceId = ServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceTopologyName, serviceName);
ServiceResource serviceResource = client.GetServiceResource(serviceResourceId);

// get the collection of this ServiceUnitResource
ServiceUnitResourceCollection collection = serviceResource.GetServiceUnitResources();

// invoke the operation
string serviceUnitName = "myServiceUnit";
ServiceUnitResourceData data = new ServiceUnitResourceData(new AzureLocation("centralus"), "myDeploymentResourceGroup", DeploymentMode.Incremental)
{
    Artifacts = new ServiceUnitArtifacts()
    {
        TemplateArtifactSourceRelativePath = "templates/myTopologyUnit.template.json",
        ParametersArtifactSourceRelativePath = "parameter/myTopologyUnit.parameters.json",
    },
    Tags =
    {
    },
};
ArmOperation<ServiceUnitResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, serviceUnitName, data);
ServiceUnitResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceUnitResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
