
import com.azure.resourcemanager.batch.models.CachingType;
import com.azure.resourcemanager.batch.models.DeploymentConfiguration;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.ImageReference;
import com.azure.resourcemanager.batch.models.ManagedDisk;
import com.azure.resourcemanager.batch.models.OSDisk;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import com.azure.resourcemanager.batch.models.StorageAccountType;
import com.azure.resourcemanager.batch.models.VirtualMachineConfiguration;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Pool Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/
     * PoolCreate_VirtualMachineConfiguration_ManagedOSDisk.json
     */
    /**
     * Sample code: CreatePool - VirtualMachineConfiguration OSDisk.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void
        createPoolVirtualMachineConfigurationOSDisk(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.pools().define("testpool").withExistingBatchAccount("default-azurebatch-japaneast", "sampleacct")
            .withVmSize("Standard_d2s_v3")
            .withDeploymentConfiguration(
                new DeploymentConfiguration().withVirtualMachineConfiguration(new VirtualMachineConfiguration()
                    .withImageReference(new ImageReference().withPublisher("microsoftwindowsserver")
                        .withOffer("windowsserver").withSku("2022-datacenter-smalldisk"))
                    .withNodeAgentSkuId("batch.node.windows amd64")
                    .withOsDisk(new OSDisk().withCaching(CachingType.READ_WRITE)
                        .withManagedDisk(new ManagedDisk().withStorageAccountType(StorageAccountType.STANDARD_SSD_LRS))
                        .withDiskSizeGB(100).withWriteAcceleratorEnabled(false))))
            .withScaleSettings(new ScaleSettings()
                .withFixedScale(new FixedScaleSettings().withTargetDedicatedNodes(1).withTargetLowPriorityNodes(0)))
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
