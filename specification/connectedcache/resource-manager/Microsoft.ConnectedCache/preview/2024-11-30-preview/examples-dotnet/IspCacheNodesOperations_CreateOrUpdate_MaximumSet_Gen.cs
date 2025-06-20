using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ConnectedCache.Models;
using Azure.ResourceManager.ConnectedCache;

// Generated from example definition: 2024-11-30-preview/IspCacheNodesOperations_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "IspCacheNodeResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IspCustomerResource created on azure
// for more information of creating IspCustomerResource, please refer to the document of IspCustomerResource
string subscriptionId = "12345678-1234-1234-1234-123456789098";
string resourceGroupName = "rgConnectedCache";
string customerResourceName = "zpqzbmanlljgmkrthtydrtneavhlnlqkdmviq";
ResourceIdentifier ispCustomerResourceId = IspCustomerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, customerResourceName);
IspCustomerResource ispCustomer = client.GetIspCustomerResource(ispCustomerResourceId);

// get the collection of this IspCacheNodeResource
IspCacheNodeCollection collection = ispCustomer.GetIspCacheNodes();

// invoke the operation
string cacheNodeResourceName = "cabakm";
IspCacheNodeData data = new IspCacheNodeData(new AzureLocation("westus"))
{
    Properties = new MccCacheNodeProperty
    {
        CacheNode = new MccCacheNodeEntity
        {
            FullyQualifiedResourceId = new ResourceIdentifier("hskxkpbiqbrbjiwdzrxndru"),
            CustomerName = "xwyqk",
            IPAddress = "voctagljcwqgcpnionqdcbjk",
            CustomerIndex = "qtoiglqaswivmkjhzogburcxtszmek",
            CacheNodeId = "xjzffjftwcgsehanoxsl",
            CacheNodeName = "mfjxb",
            CustomerAsn = 4,
            IsEnabled = true,
            MaxAllowableEgressInMbps = 29,
            IsEnterpriseManaged = true,
            CidrCsv = { "nlqlvrthafvvljuupcbcw" },
            ShouldMigrate = true,
            CidrSelectionType = 4,
        },
        AdditionalCacheNodeProperties = new MccCacheNodeAdditionalProperties
        {
            CacheNodePropertiesDetailsIssuesList = { "ex" },
            DriveConfiguration = {new CacheNodeDriveConfiguration
            {
            PhysicalPath = "/mcc",
            SizeInGb = 500,
            CacheNumber = 1,
            NginxMapping = "lijygenjq",
            }},
            BgpAsnToIPAddressMapping = "pafcimhoog",
            ProxyUri = new Uri("hplstyg"),
            OptionalProperty1 = "hvpmt",
            OptionalProperty2 = "talanelmsgxvksrzoeeontqkjzbpv",
            OptionalProperty3 = "bxkoxq",
            OptionalProperty4 = "pqlkcekupusoc",
            OptionalProperty5 = "nyvvmrjigqdufzjdvazdca",
        },
        StatusCode = "1",
        StatusText = "Success",
        StatusDetails = "djruqvptzxak",
        Error = new ResponseError(null, null),
    },
    Tags =
    {
    ["key4171"] = "qtjlszkawsdujzpgohsbw"
    },
};
ArmOperation<IspCacheNodeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, cacheNodeResourceName, data);
IspCacheNodeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IspCacheNodeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
