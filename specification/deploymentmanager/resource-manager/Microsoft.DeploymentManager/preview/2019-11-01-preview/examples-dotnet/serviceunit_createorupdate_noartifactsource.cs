using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DeploymentManager;
using Azure.ResourceManager.DeploymentManager.Models;

// Generated from example definition: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate_noartifactsource.json
// this example is just showing the usage of "ServiceUnits_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceUnitResource created on azure
// for more information of creating ServiceUnitResource, please refer to the document of ServiceUnitResource
string subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
string resourceGroupName = "myResourceGroup";
string serviceTopologyName = "myTopology";
string serviceName = "myService";
string serviceUnitName = "myServiceUnit";
ResourceIdentifier serviceUnitResourceId = ServiceUnitResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, serviceTopologyName, serviceName, serviceUnitName);
ServiceUnitResource serviceUnitResource = client.GetServiceUnitResource(serviceUnitResourceId);

// invoke the operation
ServiceUnitResourceData data = new ServiceUnitResourceData(new AzureLocation("centralus"), "myDeploymentResourceGroup", DeploymentMode.Incremental)
{
    Artifacts = new ServiceUnitArtifacts()
    {
        TemplateUri = new Uri("https://mystorageaccount.blob.core.windows.net/myartifactsource/templates/myTopologyUnit.template.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D"),
        ParametersUri = new Uri("https://mystorageaccount.blob.core.windows.net/myartifactsource/parameter/myTopologyUnit.parameters.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D"),
    },
    Tags =
    {
    },
};
ArmOperation<ServiceUnitResource> lro = await serviceUnitResource.UpdateAsync(WaitUntil.Completed, data);
ServiceUnitResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceUnitResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
