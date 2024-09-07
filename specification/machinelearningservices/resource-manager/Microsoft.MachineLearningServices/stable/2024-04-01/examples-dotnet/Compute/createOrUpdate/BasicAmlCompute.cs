using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.MachineLearning.Models;
using Azure.ResourceManager.MachineLearning;

// Generated from example definition: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/examples/Compute/createOrUpdate/BasicAmlCompute.json
// this example is just showing the usage of "Compute_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MachineLearningWorkspaceResource created on azure
// for more information of creating MachineLearningWorkspaceResource, please refer to the document of MachineLearningWorkspaceResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testrg123";
string workspaceName = "workspaces123";
ResourceIdentifier machineLearningWorkspaceResourceId = MachineLearningWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workspaceName);
MachineLearningWorkspaceResource machineLearningWorkspace = client.GetMachineLearningWorkspaceResource(machineLearningWorkspaceResourceId);

// get the collection of this MachineLearningComputeResource
MachineLearningComputeCollection collection = machineLearningWorkspace.GetMachineLearningComputes();

// invoke the operation
string computeName = "compute123";
MachineLearningComputeData data = new MachineLearningComputeData(new AzureLocation("eastus"))
{
    Properties = new AmlCompute()
    {
        Properties = new AmlComputeProperties()
        {
            OSType = MachineLearningOSType.Windows,
            VmSize = "STANDARD_NC6",
            VmPriority = MachineLearningVmPriority.Dedicated,
            VirtualMachineImageId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Compute/galleries/myImageGallery/images/myImageDefinition/versions/0.0.1",
            IsolatedNetwork = false,
            ScaleSettings = new AmlComputeScaleSettings(1)
            {
                MinNodeCount = 0,
                NodeIdleTimeBeforeScaleDown = XmlConvert.ToTimeSpan("PT5M"),
            },
            RemoteLoginPortPublicAccess = MachineLearningRemoteLoginPortPublicAccess.NotSpecified,
            EnableNodePublicIP = true,
        },
    },
};
ArmOperation<MachineLearningComputeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, computeName, data);
MachineLearningComputeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MachineLearningComputeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
