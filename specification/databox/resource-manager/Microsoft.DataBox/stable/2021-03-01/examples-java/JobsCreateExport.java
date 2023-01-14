import com.azure.resourcemanager.databox.models.AddressType;
import com.azure.resourcemanager.databox.models.ContactDetails;
import com.azure.resourcemanager.databox.models.DataAccountType;
import com.azure.resourcemanager.databox.models.DataBoxJobDetails;
import com.azure.resourcemanager.databox.models.DataExportDetails;
import com.azure.resourcemanager.databox.models.ShippingAddress;
import com.azure.resourcemanager.databox.models.Sku;
import com.azure.resourcemanager.databox.models.SkuName;
import com.azure.resourcemanager.databox.models.StorageAccountDetails;
import com.azure.resourcemanager.databox.models.TransferAllDetails;
import com.azure.resourcemanager.databox.models.TransferConfiguration;
import com.azure.resourcemanager.databox.models.TransferConfigurationTransferAllDetails;
import com.azure.resourcemanager.databox.models.TransferConfigurationType;
import com.azure.resourcemanager.databox.models.TransferType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsCreateExport.json
     */
    /**
     * Sample code: JobsCreateExport.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsCreateExport(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .jobs()
            .define("SdkJob6429")
            .withRegion("westus")
            .withExistingResourceGroup("SdkRg8091")
            .withSku(new Sku().withName(SkuName.DATA_BOX))
            .withTransferType(TransferType.EXPORT_FROM_AZURE)
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
                    .withDataExportDetails(
                        Arrays
                            .asList(
                                new DataExportDetails()
                                    .withTransferConfiguration(
                                        new TransferConfiguration()
                                            .withTransferConfigurationType(TransferConfigurationType.TRANSFER_ALL)
                                            .withTransferAllDetails(
                                                new TransferConfigurationTransferAllDetails()
                                                    .withInclude(
                                                        new TransferAllDetails()
                                                            .withDataAccountType(DataAccountType.STORAGE_ACCOUNT)
                                                            .withTransferAllBlobs(true)
                                                            .withTransferAllFiles(true))))
                                    .withAccountDetails(
                                        new StorageAccountDetails()
                                            .withStorageAccountId(
                                                "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.Storage/storageAccounts/aaaaaa2")))))
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
