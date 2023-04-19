import com.azure.resourcemanager.qumulo.models.FileSystemResource;
import java.util.HashMap;
import java.util.Map;

/** Samples for FileSystems Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/preview/2022-10-12-preview/examples/FileSystems_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: FileSystems_Update_MinimumSet_Gen.
     *
     * @param manager Entry point to QumuloManager.
     */
    public static void fileSystemsUpdateMinimumSetGen(com.azure.resourcemanager.qumulo.QumuloManager manager) {
        FileSystemResource resource =
            manager
                .fileSystems()
                .getByResourceGroupWithResponse("rgQumulo", "aaaaaaaaaaaaaaaaa", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
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
