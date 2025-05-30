
import com.azure.resourcemanager.databox.models.AddressType;
import com.azure.resourcemanager.databox.models.CreateJobValidations;
import com.azure.resourcemanager.databox.models.CreateOrderLimitForSubscriptionValidationRequest;
import com.azure.resourcemanager.databox.models.DataImportDetails;
import com.azure.resourcemanager.databox.models.DataTransferDetailsValidationRequest;
import com.azure.resourcemanager.databox.models.ModelName;
import com.azure.resourcemanager.databox.models.Preferences;
import com.azure.resourcemanager.databox.models.PreferencesValidationRequest;
import com.azure.resourcemanager.databox.models.ShippingAddress;
import com.azure.resourcemanager.databox.models.SkuAvailabilityValidationRequest;
import com.azure.resourcemanager.databox.models.SkuName;
import com.azure.resourcemanager.databox.models.StorageAccountDetails;
import com.azure.resourcemanager.databox.models.SubscriptionIsAllowedToCreateJobValidationRequest;
import com.azure.resourcemanager.databox.models.TransferType;
import com.azure.resourcemanager.databox.models.TransportPreferences;
import com.azure.resourcemanager.databox.models.TransportShipmentTypes;
import com.azure.resourcemanager.databox.models.ValidateAddress;
import java.util.Arrays;

/**
 * Samples for Service ValidateInputs.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/databox/resource-manager/Microsoft.DataBox/stable/2025-02-01/examples/ValidateInputs.json
     */
    /**
     * Sample code: ValidateInputs.
     * 
     * @param manager Entry point to DataBoxManager.
     */
    public static void validateInputs(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager.services().validateInputsWithResponse("westus", new CreateJobValidations().withIndividualRequestDetails(
            Arrays.asList(new DataTransferDetailsValidationRequest().withDataImportDetails(Arrays
                .asList(new DataImportDetails().withAccountDetails(new StorageAccountDetails().withStorageAccountId(
                    "/subscriptions/YourSubscriptionId/resourcegroups/YourResourceGroupName/providers/Microsoft.Storage/storageAccounts/YourStorageAccountName"))))
                .withDeviceType(SkuName.DATA_BOX).withTransferType(TransferType.IMPORT_TO_AZURE)
                .withModel(ModelName.DATA_BOX),
                new ValidateAddress()
                    .withShippingAddress(new ShippingAddress()
                        .withStreetAddress1("XXXX XXXX").withStreetAddress2("XXXX XXXX").withCity("XXXX XXXX")
                        .withStateOrProvince("XX").withCountry("XX").withPostalCode("fakeTokenPlaceholder")
                        .withCompanyName("XXXX XXXX").withAddressType(AddressType.COMMERCIAL))
                    .withDeviceType(SkuName.DATA_BOX)
                    .withTransportPreferences(
                        new TransportPreferences().withPreferredShipmentType(TransportShipmentTypes.MICROSOFT_MANAGED))
                    .withModel(ModelName.DATA_BOX),
                new SubscriptionIsAllowedToCreateJobValidationRequest(),
                new SkuAvailabilityValidationRequest().withDeviceType(SkuName.DATA_BOX)
                    .withTransferType(TransferType.IMPORT_TO_AZURE).withCountry("XX").withLocation("westus")
                    .withModel(ModelName.DATA_BOX),
                new CreateOrderLimitForSubscriptionValidationRequest().withDeviceType(SkuName.DATA_BOX)
                    .withModel(ModelName.DATA_BOX),
                new PreferencesValidationRequest()
                    .withPreference(new Preferences().withTransportPreferences(
                        new TransportPreferences().withPreferredShipmentType(TransportShipmentTypes.MICROSOFT_MANAGED)))
                    .withDeviceType(SkuName.DATA_BOX).withModel(ModelName.DATA_BOX))),
            com.azure.core.util.Context.NONE);
    }
}
