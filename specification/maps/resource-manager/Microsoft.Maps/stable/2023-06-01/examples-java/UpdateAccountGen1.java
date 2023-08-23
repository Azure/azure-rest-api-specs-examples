import com.azure.resourcemanager.maps.models.Kind;
import com.azure.resourcemanager.maps.models.MapsAccount;
import com.azure.resourcemanager.maps.models.Name;
import com.azure.resourcemanager.maps.models.Sku;
import java.util.HashMap;
import java.util.Map;

/** Samples for Accounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/UpdateAccountGen1.json
     */
    /**
     * Sample code: Update to Gen1 Account.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void updateToGen1Account(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        MapsAccount resource =
            manager
                .accounts()
                .getByResourceGroupWithResponse("myResourceGroup", "myMapsAccount", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withKind(Kind.GEN1).withSku(new Sku().withName(Name.S1)).apply();
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
