import com.azure.resourcemanager.hardwaresecuritymodules.models.ApiEntityReference;
import com.azure.resourcemanager.hardwaresecuritymodules.models.NetworkInterface;
import com.azure.resourcemanager.hardwaresecuritymodules.models.NetworkProfile;
import com.azure.resourcemanager.hardwaresecuritymodules.models.Sku;
import com.azure.resourcemanager.hardwaresecuritymodules.models.SkuName;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for DedicatedHsm CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/stable/2021-11-30/examples/DedicatedHsm_CreateOrUpdate.json
     */
    /**
     * Sample code: Create a new or update an existing dedicated HSM.
     *
     * @param manager Entry point to HardwareSecurityModulesManager.
     */
    public static void createANewOrUpdateAnExistingDedicatedHSM(
        com.azure.resourcemanager.hardwaresecuritymodules.HardwareSecurityModulesManager manager) {
        manager
            .dedicatedHsms()
            .define("hsm1")
            .withRegion("westus")
            .withExistingResourceGroup("hsm-group")
            .withTags(mapOf("Dept", "hsm", "Environment", "dogfood"))
            .withSku(new Sku().withName(SkuName.SAFE_NET_LUNA_NETWORK_HSM_A790))
            .withNetworkProfile(
                new NetworkProfile()
                    .withSubnet(
                        new ApiEntityReference()
                            .withId(
                                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"))
                    .withNetworkInterfaces(Arrays.asList(new NetworkInterface().withPrivateIpAddress("1.0.0.1"))))
            .withStampId("stamp01")
            .create();
    }

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
