import com.azure.resourcemanager.databox.models.CustomerResolutionCode;
import com.azure.resourcemanager.databox.models.MitigateJobRequest;
import java.util.HashMap;
import java.util.Map;

/** Samples for ResourceProvider Mitigate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobMitigate.json
     */
    /**
     * Sample code: Mitigate.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void mitigate(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .resourceProviders()
            .mitigateWithResponse(
                "TestJobName1",
                "YourResourceGroupName",
                new MitigateJobRequest()
                    .withSerialNumberCustomerResolutionMap(
                        mapOf(
                            "testDISK-1",
                            CustomerResolutionCode.MOVE_TO_CLEAN_UP_DEVICE,
                            "testDISK-2",
                            CustomerResolutionCode.RESUME)),
                com.azure.core.util.Context.NONE);
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
