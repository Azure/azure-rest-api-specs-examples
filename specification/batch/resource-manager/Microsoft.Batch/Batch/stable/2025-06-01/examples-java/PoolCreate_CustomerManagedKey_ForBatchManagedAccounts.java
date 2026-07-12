
import com.azure.resourcemanager.batch.models.BatchPoolIdentity;
import com.azure.resourcemanager.batch.models.ComputeNodeIdentityReference;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.DiskCustomerManagedKey;
import com.azure.resourcemanager.batch.models.DiskEncryptionConfiguration;
import com.azure.resourcemanager.batch.models.DiskEncryptionTarget;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.PoolIdentityType;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.UserAssignedIdentities;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Pool Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolCreate_CustomerManagedKey_ForBatchManagedAccounts.json
     */
    /**
     * Sample code: CreatePool - customer managed key for Batch managed accounts.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void
        createPoolCustomerManagedKeyForBatchManagedAccounts(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().define("testpool").withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withIdentity(new BatchPoolIdentity().withType(PoolIdentityType.USER_ASSIGNED)
                .withUserAssignedIdentities(mapOf(
                    "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
                    new UserAssignedIdentities())))
            .withVmSize("Standard_D4ds_v5")
            .withDeploymentConfiguration(
                new DeploymentConfiguration().withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2022-Datacenter").withVersion("latest"))
                    .withNodeAgentSkuId("batch.node.windows amd64")
                    .withDiskEncryptionConfiguration(new DiskEncryptionConfiguration()
                        .withTargets(Arrays.asList(DiskEncryptionTarget.OS_DISK))
                        .withCustomerManagedKey(new DiskCustomerManagedKey().withKeyUrl("fakeTokenPlaceholder")
                            .withIdentityReference(new ComputeNodeIdentityReference().withResourceId(
                                "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1"))))))
            .withScaleSettings(new ScaleSettings().withFixedScale(
                new FixedScaleSettings().withResizeTimeout(Duration.parse("PT15M")).withTargetDedicatedNodes(1)))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
