import com.azure.resourcemanager.databox.models.AddressType;
import com.azure.resourcemanager.databox.models.ShippingAddress;
import com.azure.resourcemanager.databox.models.SkuName;
import com.azure.resourcemanager.databox.models.ValidateAddress;

/** Samples for Service ValidateAddress. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2022-12-01/examples/ValidateAddressPost.json
     */
    /**
     * Sample code: ValidateAddressPost.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void validateAddressPost(com.azure.resourcemanager.databox.DataBoxManager manager) {
        manager
            .services()
            .validateAddressWithResponse(
                "westus",
                new ValidateAddress()
                    .withShippingAddress(
                        new ShippingAddress()
                            .withStreetAddress1("XXXX XXXX")
                            .withStreetAddress2("XXXX XXXX")
                            .withCity("XXXX XXXX")
                            .withStateOrProvince("XX")
                            .withCountry("XX")
                            .withPostalCode("fakeTokenPlaceholder")
                            .withCompanyName("XXXX XXXX")
                            .withAddressType(AddressType.COMMERCIAL))
                    .withDeviceType(SkuName.DATA_BOX),
                com.azure.core.util.Context.NONE);
    }
}
