import com.azure.core.util.Context;
import com.azure.resourcemanager.iothub.models.IotHubDescription;
import java.util.HashMap;
import java.util.Map;

/** Samples for IotHubResource Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_patch.json
     */
    /**
     * Sample code: IotHubResource_Update.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceUpdate(com.azure.resourcemanager.iothub.IotHubManager manager) {
        IotHubDescription resource =
            manager
                .iotHubResources()
                .getByResourceGroupWithResponse("myResourceGroup", "myHub", Context.NONE)
                .getValue();
        resource.update().withTags(mapOf("foo", "bar")).apply();
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
