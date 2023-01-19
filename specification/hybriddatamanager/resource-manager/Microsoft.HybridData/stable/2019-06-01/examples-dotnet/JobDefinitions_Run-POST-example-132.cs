using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.HybridData;
using Azure.ResourceManager.HybridData.Models;

// Generated from example definition: specification/hybriddatamanager/resource-manager/Microsoft.HybridData/stable/2019-06-01/examples/JobDefinitions_Run-POST-example-132.json
// this example is just showing the usage of "JobDefinitions_Run" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this HybridDataJobDefinitionResource created on azure
// for more information of creating HybridDataJobDefinitionResource, please refer to the document of HybridDataJobDefinitionResource
string subscriptionId = "6e0219f5-327a-4365-904f-05eed4227ad7";
string resourceGroupName = "ResourceGroupForSDKTest";
string dataManagerName = "TestAzureSDKOperations";
string dataServiceName = "DataTransformation";
string jobDefinitionName = "jobdeffromtestcode1";
ResourceIdentifier hybridDataJobDefinitionResourceId = HybridDataJobDefinitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, dataManagerName, dataServiceName, jobDefinitionName);
HybridDataJobDefinitionResource hybridDataJobDefinition = client.GetHybridDataJobDefinitionResource(hybridDataJobDefinitionResourceId);

// invoke the operation
HybridDataJobRunContent content = new HybridDataJobRunContent()
{
    UserConfirmation = UserConfirmationSetting.NotRequired,
    DataServiceInput = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["AzureStorageType"] = "Blob",
        ["BackupChoice"] = "UseExistingLatest",
        ["ContainerName"] = "containerfromtest",
        ["DeviceName"] = "8600-SHG0997877L71FC",
        ["FileNameFilter"] = "*",
        ["IsDirectoryMode"] = "false",
        ["RootDirectories"] = new object[] { "\\" },
        ["VolumeNames"] = new object[] { "TestAutomation" }
    }),
    CustomerSecrets =
    {
    },
};
await hybridDataJobDefinition.RunAsync(WaitUntil.Completed, content);

Console.WriteLine($"Succeeded");
