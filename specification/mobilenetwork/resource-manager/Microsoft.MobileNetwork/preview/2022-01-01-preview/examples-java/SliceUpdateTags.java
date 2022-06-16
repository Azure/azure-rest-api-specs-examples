import com.azure.core.util.Context;
import com.azure.resourcemanager.mobilenetwork.models.Slice;
import java.util.HashMap;
import java.util.Map;

/** Samples for Slices UpdateTags. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/SliceUpdateTags.json
     */
    /**
     * Sample code: Update mobile network slice tags.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void updateMobileNetworkSliceTags(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        Slice resource =
            manager.slices().getWithResponse("rg1", "testMobileNetwork", "testSlice", Context.NONE).getValue();
        resource.update().withTags(mapOf("tag1", "value1", "tag2", "value2")).apply();
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
