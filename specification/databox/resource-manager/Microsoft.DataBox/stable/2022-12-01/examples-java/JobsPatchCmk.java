import com.azure.resourcemanager.databox.models.JobResource;
import com.azure.resourcemanager.databox.models.KekType;
import com.azure.resourcemanager.databox.models.KeyEncryptionKey;
import com.azure.resourcemanager.databox.models.UpdateJobDetails;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/JobsPatchCmk.json
     */
    /**
     * Sample code: JobsPatchCmk.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsPatchCmk(com.azure.resourcemanager.databox.DataBoxManager manager) {
        JobResource resource =
            manager
                .jobs()
                .getByResourceGroupWithResponse(
                    "YourResourceGroupName", "TestJobName1", null, com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withDetails(
                new UpdateJobDetails()
                    .withKeyEncryptionKey(
                        new KeyEncryptionKey()
                            .withKekType(KekType.CUSTOMER_MANAGED)
                            .withKekUrl("https://xxx.xxx.xx")
                            .withKekVaultResourceId(
                                "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.KeyVault/vaults/YourKeyVaultName")))
            .apply();
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
