
import com.azure.resourcemanager.planetarycomputer.models.GeoCatalog;
import com.azure.resourcemanager.planetarycomputer.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.planetarycomputer.models.ManagedServiceIdentityUpdate;
import com.azure.resourcemanager.planetarycomputer.models.UserAssignedIdentity;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for GeoCatalogs Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-11-preview/GeoCatalogs_Update.json
     */
    /**
     * Sample code: GeoCatalogs_Update.
     * 
     * @param manager Entry point to PlanetaryComputerManager.
     */
    public static void geoCatalogsUpdate(com.azure.resourcemanager.planetarycomputer.PlanetaryComputerManager manager) {
        GeoCatalog resource = manager.geoCatalogs()
            .getByResourceGroupWithResponse("MyResourceGroup", "MyCatalog", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withTags(mapOf("MyTag", "MyValue")).withIdentity(new ManagedServiceIdentityUpdate()
            .withType(ManagedServiceIdentityType.USER_ASSIGNED)
            .withUserAssignedIdentities(mapOf(
                "/subscriptions/cd9b6cdf-dcf0-4dca-ab19-82be07b74704/resourceGroups/MyResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/MyManagedIdentity",
                new UserAssignedIdentity())))
            .apply();
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
