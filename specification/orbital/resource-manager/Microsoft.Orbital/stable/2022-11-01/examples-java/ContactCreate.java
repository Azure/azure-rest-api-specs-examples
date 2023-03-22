import com.azure.resourcemanager.orbital.models.ContactsPropertiesContactProfile;
import java.time.OffsetDateTime;

/** Samples for Contacts Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/orbital/resource-manager/Microsoft.Orbital/stable/2022-11-01/examples/ContactCreate.json
     */
    /**
     * Sample code: Create a contact.
     *
     * @param manager Entry point to OrbitalManager.
     */
    public static void createAContact(com.azure.resourcemanager.orbital.OrbitalManager manager) {
        manager
            .contacts()
            .define("contact1")
            .withExistingSpacecraft("contoso-Rgp", "CONTOSO_SAT")
            .withReservationStartTime(OffsetDateTime.parse("2023-02-22T10:58:30Z"))
            .withReservationEndTime(OffsetDateTime.parse("2023-02-22T11:10:45Z"))
            .withGroundStationName("EASTUS2_0")
            .withContactProfile(
                new ContactsPropertiesContactProfile()
                    .withId(
                        "/subscriptions/c1be1141-a7c9-4aac-9608-3c2e2f1152c3/resourceGroups/contoso-Rgp/providers/Microsoft.Orbital/contactProfiles/CONTOSO-CP"))
            .create();
    }
}
