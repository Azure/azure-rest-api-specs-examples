
import com.azure.resourcemanager.fabric.models.CapacityAdministration;
import com.azure.resourcemanager.fabric.models.FabricCapacityProperties;
import com.azure.resourcemanager.fabric.models.RpSku;
import com.azure.resourcemanager.fabric.models.RpSkuTier;
import java.util.Arrays;

/**
 * Samples for FabricCapacities CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-11-01/FabricCapacities_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update a capacity.
     * 
     * @param manager Entry point to FabricManager.
     */
    public static void createOrUpdateACapacity(com.azure.resourcemanager.fabric.FabricManager manager) {
        manager.fabricCapacities().define("azsdktest").withRegion("westcentralus").withExistingResourceGroup("TestRG")
            .withProperties(new FabricCapacityProperties().withAdministration(new CapacityAdministration()
                .withMembers(Arrays.asList("azsdktest@microsoft.com", "azsdktest2@microsoft.com"))))
            .withSku(new RpSku().withName("F2").withTier(RpSkuTier.FABRIC)).create();
    }
}
