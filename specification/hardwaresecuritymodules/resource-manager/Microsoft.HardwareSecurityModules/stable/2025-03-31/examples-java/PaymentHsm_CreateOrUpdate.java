
import com.azure.resourcemanager.hardwaresecuritymodules.models.ApiEntityReference;
import com.azure.resourcemanager.hardwaresecuritymodules.models.DedicatedHsmProperties;
import com.azure.resourcemanager.hardwaresecuritymodules.models.NetworkInterface;
import com.azure.resourcemanager.hardwaresecuritymodules.models.NetworkProfile;
import com.azure.resourcemanager.hardwaresecuritymodules.models.Sku;
import com.azure.resourcemanager.hardwaresecuritymodules.models.SkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for DedicatedHsm CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-31/PaymentHsm_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a new or update an existing payment HSM.
     * 
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void createANewOrUpdateAnExistingPaymentHSM(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager.dedicatedHsms().define("hsm1").withRegion("westus").withExistingResourceGroup("hsm-group")
            .withSku(new Sku().withName(SkuName.PAY_SHIELD10K_LMK1_CPS60))
            .withProperties(new DedicatedHsmProperties()
                .withNetworkProfile(new NetworkProfile().withSubnet(new ApiEntityReference().withResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"))
                    .withNetworkInterfaces(Arrays.asList(new NetworkInterface().withPrivateIpAddress("1.0.0.1"))))
                .withStampId("stamp01"))
            .withTags(mapOf("Dept", "hsm", "Environment", "dogfood")).create();
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
