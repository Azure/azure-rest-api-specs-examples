using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ComputeSchedule.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ComputeSchedule;

// Generated from example definition: 2025-04-15-preview/ScheduledActions_VirtualMachinesExecuteCreate_MinimumSet_Gen.json
// this example is just showing the usage of "ScheduledActions_ExecuteVirtualMachineCreateOperation" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "0505D8E4-D41A-48FB-9CA5-4AF8D93BE75F";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation locationparameter = new AzureLocation("useast");
ExecuteCreateContent content = new ExecuteCreateContent(new ResourceProvisionPayload(2)
{
    BaseProfile =
    {
    ["hardwareProfile"] = BinaryData.FromObjectAsJson(new
    {
    name = "F1",
    }),
    ["provisioningState"] = BinaryData.FromObjectAsJson(0),
    ["storageProfile"] = BinaryData.FromObjectAsJson(new
    {
    osDisk = new
    {
    osType = 0,
    },
    }),
    ["vmExtensions"] = BinaryData.FromObjectAsJson(new object[]
    {
    new
    {
    autoUpgradeMinorVersion = true,
    protectedSettings = "SomeDecryptedSecretValue",
    provisioningState = 0,
    enableAutomaticUpgrade = true,
    publisher = "Microsoft.Azure.Monitor",
    type = "AzureMonitorLinuxAgent",
    typeHandlerVersion = "1.0",
    },
    new
    {
    name = "myExtensionName",
    }
    }),
    ["resourcegroupName"] = BinaryData.FromObjectAsJson("RG5ABF491C-3164-42A6-8CB5-BF3CB53B018B"),
    ["computeApiVersion"] = BinaryData.FromObjectAsJson("2024-07-01")
    },
    ResourceOverrides = {new Dictionary<string, BinaryData>
    {
    ["name"] = BinaryData.FromObjectAsJson("myFleet_523"),
    ["location"] = BinaryData.FromObjectAsJson("LocalDev"),
    ["properties"] = BinaryData.FromObjectAsJson(new
    {
    hardwareProfile = new
    {
    vmSize = "Standard_F1s",
    },
    provisioningState = 0,
    osProfile = new
    {
    computerName = "myFleet000000",
    adminUsername = "adminUser",
    windowsConfiguration = new
    {
    additionalUnattendContent = new object[]
    {
    new
    {
    passName = "someValue",
    content = "",
    },
    new
    {
    passName = "someOtherValue",
    content = "SomeDecryptedSecretValue",
    }
    },
    },
    adminPassword = "SomeDecryptedSecretValue",
    },
    priority = 0,
    }),
    ["zones"] = BinaryData.FromObjectAsJson(new object[]
    {
    "1"
    })
    }, new Dictionary<string, BinaryData>
    {
    ["name"] = BinaryData.FromObjectAsJson("myFleet_524"),
    ["location"] = BinaryData.FromObjectAsJson("LocalDev"),
    ["properties"] = BinaryData.FromObjectAsJson(new
    {
    hardwareProfile = new
    {
    vmSize = "Standard_G1s",
    },
    provisioningState = 0,
    osProfile = new
    {
    computerName = "myFleet000000",
    adminUsername = "adminUser",
    windowsConfiguration = new
    {
    additionalUnattendContent = new object[]
    {
    new
    {
    passName = "someValue",
    content = "",
    },
    new
    {
    passName = "someOtherValue",
    content = "SomeDecryptedSecretValue",
    }
    },
    },
    adminPassword = "SomeDecryptedSecretValue",
    },
    priority = 0,
    }),
    ["zones"] = BinaryData.FromObjectAsJson(new object[]
    {
    "2"
    })
    }},
}, new ScheduledActionExecutionParameterDetail());
CreateResourceOperationResult result = await subscriptionResource.ExecuteVirtualMachineCreateOperationAsync(locationparameter, content);

Console.WriteLine($"Succeeded: {result}");
