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

// Generated from example definition: 2025-04-15-preview/ScheduledActions_VirtualMachinesExecuteDelete_MaximumSet_Gen.json
// this example is just showing the usage of "ScheduledActions_ExecuteVirtualMachineDeleteOperation" operation, for the dependent resources, they will have to be created separately.

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
AzureLocation locationparameter = new AzureLocation("east");
ExecuteDeleteContent content = new ExecuteDeleteContent(new ScheduledActionExecutionParameterDetail
{
    RetryPolicy = new UserRequestRetryPolicy
    {
        RetryCount = 2,
        RetryWindowInMinutes = 4,
    },
}, new UserRequestResources(new ResourceIdentifier[] { new ResourceIdentifier("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"), new ResourceIdentifier("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4") }))
{
    CorrelationId = "dfe927c5-16a6-40b7-a0f7-8524975ed642",
    IsForceDeletion = false,
};
DeleteResourceOperationResult result = await subscriptionResource.ExecuteVirtualMachineDeleteOperationAsync(locationparameter, content);

Console.WriteLine($"Succeeded: {result}");
