using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.SecurityCenter.Models;
using Azure.ResourceManager.SecurityCenter;

// Generated from example definition: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/PutAdaptiveApplicationControls_example.json
// this example is just showing the usage of "AdaptiveApplicationControls_Put" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityCenterLocationResource created on azure
// for more information of creating SecurityCenterLocationResource, please refer to the document of SecurityCenterLocationResource
string subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
AzureLocation ascLocation = new AzureLocation("centralus");
ResourceIdentifier securityCenterLocationResourceId = SecurityCenterLocationResource.CreateResourceIdentifier(subscriptionId, ascLocation);
SecurityCenterLocationResource securityCenterLocation = client.GetSecurityCenterLocationResource(securityCenterLocationResourceId);

// get the collection of this AdaptiveApplicationControlGroupResource
AdaptiveApplicationControlGroupCollection collection = securityCenterLocation.GetAdaptiveApplicationControlGroups();

// invoke the operation
string groupName = "ERELGROUP1";
AdaptiveApplicationControlGroupData data = new AdaptiveApplicationControlGroupData
{
    EnforcementMode = AdaptiveApplicationControlEnforcementMode.Audit,
    ProtectionMode = new SecurityCenterFileProtectionMode
    {
        Exe = AdaptiveApplicationControlEnforcementMode.Audit,
        Msi = AdaptiveApplicationControlEnforcementMode.None,
        Script = AdaptiveApplicationControlEnforcementMode.None,
    },
    VmRecommendations = {new VmRecommendation
    {
    ConfigurationStatus = SecurityCenterConfigurationStatus.Configured,
    RecommendationAction = RecommendationAction.Recommended,
    ResourceId = new ResourceIdentifier("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/erelh-stable/providers/microsoft.compute/virtualmachines/erelh-16090"),
    EnforcementSupport = SecurityCenterVmEnforcementSupportState.Supported,
    }, new VmRecommendation
    {
    ConfigurationStatus = SecurityCenterConfigurationStatus.Configured,
    RecommendationAction = RecommendationAction.Recommended,
    ResourceId = new ResourceIdentifier("/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourcegroups/matanvs/providers/microsoft.compute/virtualmachines/matanvs19"),
    EnforcementSupport = SecurityCenterVmEnforcementSupportState.Supported,
    }},
    PathRecommendations = {new PathRecommendation
    {
    Path = "[Exe] O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US\\*\\*\\0.0.0.0",
    Action = RecommendationAction.Recommended,
    IotSecurityRecommendationType = new IotSecurityRecommendationType("PublisherSignature"),
    PublisherInfo = new SecurityCenterPublisherInfo
    {
    PublisherName = "O=MICROSOFT CORPORATION, L=REDMOND, S=WASHINGTON, C=US",
    ProductName = "*",
    BinaryName = "*",
    Version = "0.0.0.0",
    },
    IsCommon = true,
    UserSids = {"S-1-1-0"},
    Usernames = {new UserRecommendation
    {
    Username = "Everyone",
    RecommendationAction = RecommendationAction.Recommended,
    }},
    FileType = PathRecommendationFileType.Exe,
    ConfigurationStatus = SecurityCenterConfigurationStatus.Configured,
    }, new PathRecommendation
    {
    Path = "%OSDRIVE%\\WINDOWSAZURE\\SECAGENT\\WASECAGENTPROV.EXE",
    Action = RecommendationAction.Recommended,
    IotSecurityRecommendationType = new IotSecurityRecommendationType("ProductSignature"),
    PublisherInfo = new SecurityCenterPublisherInfo
    {
    PublisherName = "CN=MICROSOFT AZURE DEPENDENCY CODE SIGN",
    ProductName = "MICROSOFT® COREXT",
    BinaryName = "*",
    Version = "0.0.0.0",
    },
    IsCommon = true,
    UserSids = {"S-1-1-0"},
    Usernames = {new UserRecommendation
    {
    Username = "NT AUTHORITY\\SYSTEM",
    RecommendationAction = RecommendationAction.Recommended,
    }},
    FileType = PathRecommendationFileType.Exe,
    ConfigurationStatus = SecurityCenterConfigurationStatus.Configured,
    }, new PathRecommendation
    {
    Path = "%OSDRIVE%\\WINDOWSAZURE\\PACKAGES_201973_7415\\COLLECTGUESTLOGS.EXE",
    Action = RecommendationAction.Recommended,
    IotSecurityRecommendationType = new IotSecurityRecommendationType("PublisherSignature"),
    PublisherInfo = new SecurityCenterPublisherInfo
    {
    PublisherName = "CN=MICROSOFT AZURE DEPENDENCY CODE SIGN",
    ProductName = "*",
    BinaryName = "*",
    Version = "0.0.0.0",
    },
    IsCommon = true,
    UserSids = {"S-1-1-0"},
    Usernames = {new UserRecommendation
    {
    Username = "NT AUTHORITY\\SYSTEM",
    RecommendationAction = RecommendationAction.Recommended,
    }},
    FileType = PathRecommendationFileType.Exe,
    ConfigurationStatus = SecurityCenterConfigurationStatus.Configured,
    }, new PathRecommendation
    {
    Path = "C:\\directory\\file.exe",
    Action = RecommendationAction.Add,
    IotSecurityRecommendationType = new IotSecurityRecommendationType("File"),
    IsCommon = true,
    }},
};
ArmOperation<AdaptiveApplicationControlGroupResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, groupName, data);
AdaptiveApplicationControlGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AdaptiveApplicationControlGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
