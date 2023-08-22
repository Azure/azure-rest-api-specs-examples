import com.azure.resourcemanager.maps.models.Kind;
import com.azure.resourcemanager.maps.models.LinkedResource;
import com.azure.resourcemanager.maps.models.ManagedServiceIdentity;
import com.azure.resourcemanager.maps.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.maps.models.MapsAccount;
import com.azure.resourcemanager.maps.models.Name;
import com.azure.resourcemanager.maps.models.Sku;
import com.azure.resourcemanager.maps.models.UserAssignedIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Accounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/UpdateAccountManagedIdentity.json
     */
    /**
     * Sample code: Update Account Managed Identities.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void updateAccountManagedIdentities(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        MapsAccount resource =
            manager
                .accounts()
                .getByResourceGroupWithResponse("myResourceGroup", "myMapsAccount", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withKind(Kind.GEN2)
            .withSku(new Sku().withName(Name.G2))
            .withIdentity(
                new ManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName",
                            new UserAssignedIdentity())))
            .withLinkedResources(
                Arrays
                    .asList(
                        new LinkedResource()
                            .withUniqueName("myBatchStorageAccount")
                            .withId(
                                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/accounts/{storageName}")))
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
