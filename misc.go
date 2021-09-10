package helper

import (
    "encoding/json"
)

func Unmarshal(data interface{}, v interface{}) error {
    switch d := data.(type) {
    case []byte:
        return json.Unmarshal(d, v)
    default:
        if b, err := json.Marshal(&data); err != nil {
            return err
        } else {
            return json.Unmarshal(b, v)
        }
    }
}
