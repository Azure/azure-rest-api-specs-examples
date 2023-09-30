using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.MachineLearning;
using Azure.ResourceManager.MachineLearning.Models;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2023-06-01-preview/examples/WorkspaceConnection/update.json
// this example is just showing the usage of "WorkspaceConnections_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceConnectionResource created on azure
// for more information of creating MachineLearningWorkspaceConnectionResource, please refer to the document of MachineLearningWorkspaceConnectionResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "test-rg";
string workspaceName = "aml-workspace-name";
string connectionName = "some_string";
ResourceIdentifier machineLearningWorkspaceConnectionResourceId = MachineLearningWorkspaceConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName, connectionName);
MachineLearningWorkspaceConnectionResource machineLearningWorkspaceConnection = client.GetMachineLearningWorkspaceConnectionResource(machineLearningWorkspaceConnectionResourceId);

// invoke the operation
MachineLearningWorkspaceConnectionPatch patch = new MachineLearningWorkspaceConnectionPatch()
{
    Properties = new AccessKeyAuthTypeWorkspaceConnectionProperties()
    {
        Credentials = new WorkspaceConnectionAccessKey()
        {
            AccessKeyId = "some_string",
            SecretAccessKey = "some_string",
        },
        Category = MachineLearningConnectionCategory.AdlsGen2,
        ExpiryOn = DateTimeOffset.Parse("2020-01-01T00:00:00Z"),
        Metadata = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
        {
        }),
        Target = "some_string",
    },
};
MachineLearningWorkspaceConnectionResource result = await machineLearningWorkspaceConnection.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningWorkspaceConnectionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
