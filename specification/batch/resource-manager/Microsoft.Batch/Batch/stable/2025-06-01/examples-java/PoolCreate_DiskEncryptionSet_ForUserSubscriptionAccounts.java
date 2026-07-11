
import com.azure.resourcemanager.batch.models.ComputeNodeFillType;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.DiskEncryptionSetParameters;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.ManagedDisk;
import com.azure.resourcemanager.batch.models.OSDisk;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.SecurityProfile;
import com.azure.resourcemanager.batch.models.StorageAccountType;
import com.azure.resourcemanager.batch.models.TaskSchedulingPolicy;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.time.Duration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Pool Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/PoolCreate_DiskEncryptionSet_ForUserSubscriptionAccounts.json
     */
    /**
     * Sample code: CreatePool - disk encryption set for user subscription accounts.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void
        createPoolDiskEncryptionSetForUserSubscriptionAccounts(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().define("testpool").withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withVmSize("Standard_D4ds_v5")
            .withDeploymentConfiguration(
                new DeploymentConfiguration().withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                    .withImageReference(new ImageReference().withPublisher("MicrosoftWindowsServer")
                        .withOffer("WindowsServer").withSku("2022-Datacenter").withVersion("latest"))
                    .withNodeAgentSkuId("batch.node.windows amd64")
                    .withOsDisk(new OSDisk().withManagedDisk(new ManagedDisk()
                        .withStorageAccountType(StorageAccountType.STANDARD_LRS)
                        .withDiskEncryptionSet(new DiskEncryptionSetParameters().withId(
                            "/subscriptions/12345678-1234-1234-1234-123456789012/resourceGroups/default-azurebatch-japaneast/providers/Microsoft.Compute/diskEncryptionSets/DiskEncryptionSetId"))))
                    .withSecurityProfile(new SecurityProfile().withEncryptionAtHost(false))))
            .withScaleSettings(new ScaleSettings().withFixedScale(
                new FixedScaleSettings().withResizeTimeout(Duration.parse("PT15M")).withTargetDedicatedNodes(1)))
            .withTaskSchedulingPolicy(new TaskSchedulingPolicy().withNodeFillType(ComputeNodeFillType.PACK)).create();
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
