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

// Generated from example definition: 2025-04-15-preview/ScheduledActions_VirtualMachinesSubmitHibernate_MaximumSet_Gen.json
// this example is just showing the usage of "ScheduledActions_SubmitVirtualMachineHibernate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "CB26D7CB-3E27-465F-99C8-EAF7A4118245";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation locationparameter = new AzureLocation("rhadyapnyvmobwg");
SubmitHibernateContent content = new SubmitHibernateContent(new UserRequestSchedule(ScheduledActionDeadlineType.Unknown)
{
    UserRequestDeadline = DateTimeOffset.Parse("2025-04-17T00:23:56.803Z"),
    UserRequestTimezone = "aigbjdnldtzkteqi",
}, new ScheduledActionExecutionParameterDetail
{
    OptimizationPreference = ScheduledActionOptimizationPreference.Cost,
    RetryPolicy = new UserRequestRetryPolicy
    {
        RetryCount = 17,
        RetryWindowInMinutes = 29,
    },
}, new UserRequestResources(new ResourceIdentifier[] { new ResourceIdentifier("/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource4") }), "b211f086-4b91-4686-a453-2f5c012e4d80");
HibernateResourceOperationResult result = await subscriptionResource.SubmitVirtualMachineHibernateAsync(locationparameter, content);

Console.WriteLine($"Succeeded: {result}");
