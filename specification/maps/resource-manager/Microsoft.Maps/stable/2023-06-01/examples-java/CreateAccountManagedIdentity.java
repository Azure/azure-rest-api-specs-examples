import com.azure.resourcemanager.maps.fluent.models.MapsAccountProperties;
import com.azure.resourcemanager.maps.models.Kind;
import com.azure.resourcemanager.maps.models.LinkedResource;
import com.azure.resourcemanager.maps.models.ManagedServiceIdentity;
import com.azure.resourcemanager.maps.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.maps.models.Name;
import com.azure.resourcemanager.maps.models.Sku;
import com.azure.resourcemanager.maps.models.UserAssignedIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Accounts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/CreateAccountManagedIdentity.json
     */
    /**
     * Sample code: Create Account with Managed Identities.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void createAccountWithManagedIdentities(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager
            .accounts()
            .define("myMapsAccount")
            .withRegion("eastus")
            .withExistingResourceGroup("myResourceGroup")
            .withSku(new Sku().withName(Name.G2))
            .withTags(mapOf("test", "true"))
            .withKind(Kind.GEN2)
            .withIdentity(
                new ManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED_USER_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName",
                            new UserAssignedIdentity())))
            .withProperties(
                new MapsAccountProperties()
                    .withDisableLocalAuth(false)
                    .withLinkedResources(
                        Arrays
                            .asList(
                                new LinkedResource()
                                    .withUniqueName("myBatchStorageAccount")
                                    .withId(
                                        "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc"),
                                new LinkedResource()
                                    .withUniqueName("myBlobDataSource")
                                    .withId(
                                        "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc"))))
            .create();
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
