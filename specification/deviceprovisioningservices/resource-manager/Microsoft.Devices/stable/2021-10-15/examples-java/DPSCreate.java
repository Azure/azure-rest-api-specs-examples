import com.azure.resourcemanager.deviceprovisioningservices.models.IotDpsPropertiesDescription;
import com.azure.resourcemanager.deviceprovisioningservices.models.IotDpsSku;
import com.azure.resourcemanager.deviceprovisioningservices.models.IotDpsSkuInfo;
import java.util.HashMap;
import java.util.Map;

/** Samples for IotDpsResource CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSCreate.json
     */
    /**
     * Sample code: DPSCreate.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSCreate(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .define("myFirstProvisioningService")
            .withRegion("East US")
            .withExistingResourceGroup("myResourceGroup")
            .withProperties(new IotDpsPropertiesDescription().withEnableDataResidency(false))
            .withSku(new IotDpsSkuInfo().withName(IotDpsSku.S1).withCapacity(1L))
            .withTags(mapOf())
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
