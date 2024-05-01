package main

type ListVariantResponse struct {
	Data []struct {
		ID         int    `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Equipment       int    `json:"equipment"`
			Instances       []int  `json:"instances"`
			AttributeValues []int  `json:"attributeValues"`
			Description     any    `json:"description"`
			InstanceBuffer  int    `json:"instanceBuffer"`
			State           string `json:"state"`
			CreatedAt       string `json:"createdAt"`
			CreatedBy       string `json:"createdBy"`
			ModifiedAt      string `json:"modifiedAt"`
			ModifiedBy      string `json:"modifiedBy"`
			RemovedAt       any    `json:"removedAt"`
			RemovedBy       any    `json:"removedBy"`
		} `json:"attributes"`
		Relationships struct {
			AttributeValues struct {
				Data []struct {
					ID         int    `json:"id"`
					Type       string `json:"type"`
					Attributes struct {
						Attribute      int    `json:"attribute"`
						AttributeLabel string `json:"attributeLabel"`
						Value          string `json:"value"`
					} `json:"attributes"`
					Relationships any `json:"relationships"`
				} `json:"data"`
			} `json:"attributeValues"`
			Instances struct {
				Data []struct {
					ID         int    `json:"id"`
					Type       string `json:"type"`
					Attributes struct {
						Location           int    `json:"location"`
						Equipment          int    `json:"equipment"`
						Variant            int    `json:"variant"`
						Identifier         string `json:"identifier"`
						LabelIdentifier    any    `json:"labelIdentifier"`
						Status             string `json:"status"`
						ReservationEnabled bool   `json:"reservationEnabled"`
						Notes              any    `json:"notes"`
						ReturnAt           any    `json:"returnAt"`
						LastLoanID         int    `json:"lastLoanId"`
						State              string `json:"state"`
						CreatedAt          string `json:"createdAt"`
						CreatedBy          string `json:"createdBy"`
						ModifiedAt         string `json:"modifiedAt"`
						ModifiedBy         string `json:"modifiedBy"`
						RemovedAt          any    `json:"removedAt"`
						RemovedBy          any    `json:"removedBy"`
					} `json:"attributes"`
					Relationships any `json:"relationships"`
				} `json:"data"`
			} `json:"instances"`
			Equipment struct {
				Data struct {
					ID         int    `json:"id"`
					Type       string `json:"type"`
					Attributes struct {
						Location        int    `json:"location"`
						Category        int    `json:"category"`
						Variants        []int  `json:"variants"`
						AttributeValues []any  `json:"attributeValues"`
						Images          []any  `json:"images"`
						Brand           any    `json:"brand"`
						Name            string `json:"name"`
						Slug            string `json:"slug"`
						Description     any    `json:"description"`
						IsPublic        bool   `json:"isPublic"`
						IsHidden        bool   `json:"isHidden"`
						State           string `json:"state"`
						CreatedAt       string `json:"createdAt"`
						CreatedBy       string `json:"createdBy"`
						ModifiedAt      string `json:"modifiedAt"`
						ModifiedBy      string `json:"modifiedBy"`
						RemovedAt       any    `json:"removedAt"`
						RemovedBy       any    `json:"removedBy"`
					} `json:"attributes"`
					Relationships any `json:"relationships"`
				} `json:"data"`
			} `json:"equipment"`
		} `json:"relationships"`
	} `json:"data"`
	Meta struct {
		Total   int `json:"total"`
		Limit   int `json:"limit"`
		Offset  int `json:"offset"`
		Filters struct {
			State     string `json:"state"`
			Equipment int    `json:"equipment"`
		} `json:"filters"`
	} `json:"meta"`
}
