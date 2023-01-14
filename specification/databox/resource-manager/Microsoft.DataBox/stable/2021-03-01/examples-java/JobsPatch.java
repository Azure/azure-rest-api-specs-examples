import com.azure.resourcemanager.databox.models.AddressType;
import com.azure.resourcemanager.databox.models.ContactDetails;
import com.azure.resourcemanager.databox.models.JobResource;
import com.azure.resourcemanager.databox.models.ShippingAddress;
import com.azure.resourcemanager.databox.models.UpdateJobDetails;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Jobs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/stable/2021-03-01/examples/JobsPatch.json
     */
    /**
     * Sample code: JobsPatch.
     *
     * @param manager Entry point to DataBoxManager.
     */
    public static void jobsPatch(com.azure.resourcemanager.databox.DataBoxManager manager) {
        JobResource resource =
            manager
                .jobs()
                .getByResourceGroupWithResponse("SdkRg5154", "SdkJob952", "details", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withDetails(
                new UpdateJobDetails()
                    .withContactDetails(
                        new ContactDetails()
                            .withContactName("Update Job")
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
                            .withAddressType(AddressType.COMMERCIAL)))
            .apply();
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
