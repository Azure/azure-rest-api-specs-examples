
import com.azure.resourcemanager.databox.models.AddressType;
import com.azure.resourcemanager.databox.models.ContactDetails;
import com.azure.resourcemanager.databox.models.DataBoxJobDetails;
import com.azure.resourcemanager.databox.models.DataImportDetails;
import com.azure.resourcemanager.databox.models.ResourceIdentity;
import com.azure.resourcemanager.databox.models.ShippingAddress;
import com.azure.resourcemanager.databox.models.Sku;
import com.azure.resourcemanager.databox.models.SkuName;
import com.azure.resourcemanager.databox.models.StorageAccountDetails;
import com.azure.resourcemanager.databox.models.TransferType;
import com.azure.resourcemanager.databox.models.UserAssignedIdentity;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Jobs Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/
     * JobsCreateWithUserAssignedIdentity.json
     */
    /**
     * Sample code: JobsCreateWithUserAssignedIdentity.
     * 
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsCreateWithUserAssignedIdentity(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager.jobs().define("TestJobName1").withRegion("westus").withExistingResourceGroup("YourResourceGroupName")
            .withSku(new Sku().withName(SkuName.DATA_BOX)).withTransferType(TransferType.IMPORT_TO_AZURE)
            .withIdentity(new ResourceIdentity().withType("UserAssigned").withUserAssignedIdentities(mapOf(
                "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testIdentity",
                new UserAssignedIdentity())))
            .withDetails(new DataBoxJobDetails()
                .withContactDetails(new ContactDetails().withContactName("XXXX XXXX").withPhone("0000000000")
                    .withPhoneExtension("").withEmailList(Arrays.asList("xxxx@xxxx.xxx")))
                .withShippingAddress(new ShippingAddress().withStreetAddress1("XXXX XXXX")
                    .withStreetAddress2("XXXX XXXX").withCity("XXXX XXXX").withStateOrProvince("XX").withCountry("XX")
                    .withPostalCode("fakeTokenPlaceholder").withCompanyName("XXXX XXXX")
                    .withAddressType(AddressType.COMMERCIAL))
                .withDataImportDetails(Arrays
                    .asList(new DataImportDetails().withAccountDetails(new StorageAccountDetails().withStorageAccountId(
                        "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName")))))
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
