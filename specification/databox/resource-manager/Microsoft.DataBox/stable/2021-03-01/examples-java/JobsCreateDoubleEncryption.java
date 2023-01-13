import com.azure.resourcemanager.databox.models.AddressType;
import com.azure.resourcemanager.databox.models.ContactDetails;
import com.azure.resourcemanager.databox.models.DataBoxJobDetails;
import com.azure.resourcemanager.databox.models.DataImportDetails;
import com.azure.resourcemanager.databox.models.DoubleEncryption;
import com.azure.resourcemanager.databox.models.EncryptionPreferences;
import com.azure.resourcemanager.databox.models.Preferences;
import com.azure.resourcemanager.databox.models.ShippingAddress;
import com.azure.resourcemanager.databox.models.Sku;
import com.azure.resourcemanager.databox.models.SkuName;
import com.azure.resourcemanager.databox.models.StorageAccountDetails;
import com.azure.resourcemanager.databox.models.TransferType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsCreateDoubleEncryption.json
     */
    /**
     * Sample code: JobsCreateDoubleEncryption.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsCreateDoubleEncryption(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .define("SdkJob6599")
            .withRegion("westus")
            .withExistingResourceGroup("SdkRg608")
            .withSku(new Sku().withName(SkuName.DATA_BOX))
            .withTransferType(TransferType.IMPORT_TO_AZURE)
            .withDetails(
                new DataBoxJobDetails()
                    .withContactDetails(
                        new ContactDetails()
                            .withContactName("Public SDK Test")
                            .withPhone("1234567890")
                            .withPhoneExtension("1234")
                            .withEmailList(Arrays.asList("testing@microsoft.com")))
                    .withShippingAddress(
                        new ShippingAddress()
                            .withStreetAddress1("16 TOWNSEND ST")
                            .withStreetAddress2("Unit 1")
                            .withCity("San Francisco")
                            .withStateOrProvince("CA")
                            .withCountry("US")
                            .withPostalCode("fakeTokenPlaceholder")
                            .withCompanyName("Microsoft")
                            .withAddressType(AddressType.COMMERCIAL))
                    .withDataImportDetails(
                        Arrays
                            .asList(
                                new DataImportDetails()
                                    .withAccountDetails(
                                        new StorageAccountDetails()
                                            .withStorageAccountId(
                                                "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourcegroups/databoxbvt/providers/Microsoft.Storage/storageAccounts/databoxbvttestaccount"))))
                    .withPreferences(
                        new Preferences()
                            .withEncryptionPreferences(
                                new EncryptionPreferences().withDoubleEncryption(DoubleEncryption.ENABLED))))
            .create();
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
