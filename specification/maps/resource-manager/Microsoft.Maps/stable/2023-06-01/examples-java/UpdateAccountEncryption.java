import com.azure.resourcemanager.maps.models.CustomerManagedKeyEncryption;
import com.azure.resourcemanager.maps.models.CustomerManagedKeyEncryptionKeyIdentity;
import com.azure.resourcemanager.maps.models.Encryption;
import com.azure.resourcemanager.maps.models.IdentityType;
import com.azure.resourcemanager.maps.models.ManagedServiceIdentity;
import com.azure.resourcemanager.maps.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.maps.models.MapsAccount;
import java.util.HashMap;
import java.util.Map;

/** Samples for Accounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/UpdateAccountEncryption.json
     */
    /**
     * Sample code: Update Account Encryption.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void updateAccountEncryption(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        MapsAccount resource =
            manager
                .accounts()
                .getByResourceGroupWithResponse("myResourceGroup", "myMapsAccount", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withIdentity(
                new ManagedServiceIdentity()
                    .withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED)
                    .withUserAssignedIdentities(
                        mapOf(
                            "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName",
                            null)))
            .withEncryption(
                new Encryption()
                    .withCustomerManagedKeyEncryption(
                        new CustomerManagedKeyEncryption()
                            .withKeyEncryptionKeyIdentity(
                                new CustomerManagedKeyEncryptionKeyIdentity()
                                    .withIdentityType(IdentityType.SYSTEM_ASSIGNED_IDENTITY))
                            .withKeyEncryptionKeyUrl("fakeTokenPlaceholder")))
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
