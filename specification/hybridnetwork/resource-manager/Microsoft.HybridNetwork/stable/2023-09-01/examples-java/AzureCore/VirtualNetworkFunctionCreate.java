
import com.azure.resourcemanager.hybridnetwork.models.NetworkFunctionValueWithoutSecrets;
import com.azure.resourcemanager.hybridnetwork.models.NfviType;
import com.azure.resourcemanager.hybridnetwork.models.OpenDeploymentResourceReference;

/**
 * Samples for NetworkFunctions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/AzureCore/
     * VirtualNetworkFunctionCreate.json
     */
    /**
     * Sample code: Create virtual network function resource on AzureCore.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createVirtualNetworkFunctionResourceOnAzureCore(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctions().define("testNf").withRegion("eastus").withExistingResourceGroup("rg")
            .withProperties(new NetworkFunctionValueWithoutSecrets()
                .withNetworkFunctionDefinitionVersionResourceReference(new OpenDeploymentResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/publishers/testVendor/networkFunctionDefinitionGroups/testnetworkFunctionDefinitionGroupName/networkFunctionDefinitionVersions/1.0.1"))
                .withNfviType(NfviType.AZURE_CORE).withNfviId("/subscriptions/subid/resourceGroups/testResourceGroup")
                .withAllowSoftwareUpdate(false).withDeploymentValues(
                    "{\"virtualMachineName\":\"test-VM\",\"cpuCores\":4,\"memorySizeGB\":8,\"cloudServicesNetworkAttachment\":{\"attachedNetworkId\":\"test-csnet\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-cs-vlan\"},\"networkAttachments\":[{\"attachedNetworkId\":\"test-l3vlan\",\"defaultGateway\":\"True\",\"ipAllocationMethod\":\"Dynamic\",\"networkAttachmentName\":\"test-vlan\"}],\"sshPublicKeys\":[{\"keyData\":\"ssh-rsa CMIIIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA0TqlveKKlc2MFvEmuXJiLGBsY1t4ML4uiRADGSZlnc+7Ugv3h+MCjkkwOKiOdsNo8k4KSBIG5GcQfKYOOd17AJvqCL6cGQbaLuqv0a64jeDm8oO8/xN/IM0oKw7rMr/2oAJOgIsfeXPkRxWWic9AVIS++H5Qi2r7bUFX+cqFsyUCAwEBBQ==\"}],\"storageProfile\":{\"osDisk\":{\"createOption\":\"Ephemeral\",\"deleteOption\":\"Delete\",\"diskSizeGB\":10}},\"userData\":\"testUserData\",\"adminUsername\":\"testUser\",\"virtioInterface\":\"Transitional\",\"isolateEmulatorThread\":\"False\",\"bootMethod\":\"BIOS\",\"placementHints\":[]}"))
            .create();
    }
}
