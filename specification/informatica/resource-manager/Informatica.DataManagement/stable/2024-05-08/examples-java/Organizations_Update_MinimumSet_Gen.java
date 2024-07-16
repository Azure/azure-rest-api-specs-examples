
import com.azure.resourcemanager.informaticadatamanagement.models.InformaticaOrganizationResource;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Organizations Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/informatica/resource-manager/Informatica.DataManagement/stable/2024-05-08/examples/
     * Organizations_Update_MinimumSet_Gen.json
     */
    /**
     * Sample code: Organizations_Update_Min.
     * 
     * @param manager Entry point to InformaticaDataManagementManager.
     */
    public static void organizationsUpdateMin(
        com.azure.resourcemanager.informaticadatamanagement.InformaticaDataManagementManager manager) {
        InformaticaOrganizationResource resource = manager.organizations()
            .getByResourceGroupWithResponse("rgopenapi", "-", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
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
