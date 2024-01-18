using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.SecurityCenter;
using Azure.ResourceManager.SecurityCenter.Models;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/preview/2023-10-01-preview/examples/SecurityConnectors/PatchSecurityConnector_example.json
// this example is just showing the usage of "SecurityConnectors_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityConnectorResource created on azure
// for more information of creating SecurityConnectorResource, please refer to the document of SecurityConnectorResource
string subscriptionId = "a5caac9c-5c04-49af-b3d0-e204f40345d5";
string resourceGroupName = "exampleResourceGroup";
string securityConnectorName = "exampleSecurityConnectorName";
ResourceIdentifier securityConnectorResourceId = SecurityConnectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, securityConnectorName);
SecurityConnectorResource securityConnector = client.GetSecurityConnectorResource(securityConnectorResourceId);

// invoke the operation
SecurityConnectorData data = new SecurityConnectorData(new AzureLocation("Central US"))
{
    HierarchyIdentifier = "exampleHierarchyId",
    EnvironmentName = SecurityCenterCloudName.Aws,
    Offerings =
    {
    new CspmMonitorAwsOffering()
    {
    CloudRoleArn = "arn:aws:iam::00000000:role/ASCMonitor",
    }
    },
    EnvironmentData = new AwsEnvironment(),
    ETag = new ETag("etag value (must be supplied for update)"),
    Tags =
    {
    },
};
SecurityConnectorResource result = await securityConnector.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SecurityConnectorData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
