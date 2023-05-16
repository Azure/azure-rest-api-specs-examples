import com.azure.resourcemanager.appcontainers.models.ManagedCertificate;
import java.util.HashMap;
import java.util.Map;

/** Samples for ManagedCertificates Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ManagedCertificates_Patch.json
     */
    /**
     * Sample code: Patch Managed Certificate.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void patchManagedCertificate(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        ManagedCertificate resource =
            manager
                .managedCertificates()
                .getWithResponse(
                    "examplerg", "testcontainerenv", "certificate-firendly-name", com.azure.core.util.Context.NONE)
                .getValue();
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
